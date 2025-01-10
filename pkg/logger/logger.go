package logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ctxKey string

const (
	CorrelationIDKey ctxKey = "correlation_id"
	RequestIDKey     ctxKey = "request_id"
	ServiceKey       ctxKey = "service"
)

// Logger is the interface for logging
type Logger struct {
	*zap.Logger
	serviceName string
}

type Config struct {
	Level       string
	ServiceName string
	Environment string
}

// New creates a new logger

func NewLogger(cfg *LogConfig) (*Logger, error) {
	if cfg.LogDir == "" {
		cfg.LogDir = filepath.Join("logs", cfg.Environment, cfg.ServiceName,
			time.Now().Format("2006-01-02"))
	}

	if err := os.MkdirAll(cfg.LogDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %w", err)
	}

	logFile := filepath.Join(cfg.LogDir, fmt.Sprintf("%s.log", cfg.ServiceName))
	errorFile := filepath.Join(cfg.LogDir, fmt.Sprintf("%s.error.log", cfg.ServiceName))

	accessLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	errorLogger := &lumberjack.Logger{
		Filename:   errorFile,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}

	level := getLogLevel(cfg.Level)

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(createEncoderConfig()),
			zapcore.AddSync(accessLogger),
			lowPriority,
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(createEncoderConfig()),
			zapcore.AddSync(errorLogger),
			highPriority,
		),
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(createEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			level,
		),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &Logger{
		Logger:      logger,
		serviceName: cfg.ServiceName,
	}, nil
}

func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

// with context

func (l *Logger) WithContext(ctx context.Context) *zap.Logger {
	var fields []zap.Field

	// Add correlation Id if exists

	if cid, ok := ctx.Value(CorrelationIDKey).(string); ok {
		fields = append(fields, zap.String("correlation_id", cid))
	}

	// Add request Id if exists

	if rid, ok := ctx.Value(RequestIDKey).(string); ok {
		fields = append(fields, zap.String("request_id", rid))
	}

	return l.With(fields...)
}

// WithField adds a field to the logger
func (l *Logger) WithField(key string, value interface{}) *zap.Logger {
	return l.With(zap.Any(key, value))
}

// WithError adds an error to the logger
func (l *Logger) WithError(err error) *zap.Logger {
	return l.With(zap.Error(err))
}

// newContext

func (l *Logger) NewContext(ctx context.Context) context.Context {
	cid := uuid.New().String()
	return context.WithValue(ctx, CorrelationIDKey, cid)
}

func (l *Logger) RequestStarted(ctx context.Context, method, path string) {
	l.WithContext(ctx).Info("request_started",
		zap.String("method", method),
		zap.String("path", path),
		zap.Time("start_time", time.Now()),
	)
}

func (l *Logger) RequestCompleted(ctx context.Context, method, path string, statusCode int, duration time.Duration) {
	l.WithContext(ctx).Info("request_completed",
		zap.String("method", method),
		zap.String("path", path),
		zap.Int("status_code", statusCode),
		zap.Duration("duration", duration),
	)
}

func (l *Logger) Error(ctx context.Context, msg string, err error, fields ...zap.Field) {
	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	l.WithContext(ctx).Error(msg, fields...)
}

func createEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
