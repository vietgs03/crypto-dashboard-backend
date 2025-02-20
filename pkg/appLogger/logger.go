package appLogger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"crypto-dashboard/pkg/constants"
	"crypto-dashboard/pkg/response"
	"crypto-dashboard/pkg/settings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is the interface for logging
type (
	Logger struct {
		*zap.Logger
		serviceName string
	}
)

func NewLogger(cfg *settings.LoggerSetting, serverConfig *settings.ServerSetting) (*Logger, *response.AppError) {
	if cfg.LogDir == "" {
		cfg.LogDir = filepath.Join("logs", serverConfig.Environment, serverConfig.ServiceName,
			time.Now().Format("2006-01-02"))
	}

	if err := os.MkdirAll(cfg.LogDir, 0o755); err != nil {
		return nil, response.ServerError("failed to create log directory")
	}

	logFile := filepath.Join(cfg.LogDir, fmt.Sprintf("%s.log", serverConfig.ServiceName))
	errorFile := filepath.Join(cfg.LogDir, fmt.Sprintf("%s.error.log", serverConfig.ServiceName))

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
		serviceName: serverConfig.ServiceName,
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

	if cid, ok := ctx.Value(constants.CORRELATION_ID_KEY).(string); ok {
		fields = append(fields, zap.String("cid", cid))
	}

	// Add request Id if exists

	if rid, ok := ctx.Value(constants.REQUEST_ID_KEY).(string); ok {
		fields = append(fields, zap.String("rid", rid))
	}

	return l.With(fields...)
}

// WithField adds a field to the logger
func (l *Logger) WithField(key string, value any) *zap.Logger {
	return l.With(zap.Any(key, value))
}

// WithError adds an error to the logger
func (l *Logger) WithError(err error) *zap.Logger {
	return l.With(zap.Error(err))
}

func (l *Logger) ErrorWithCtx(ctx context.Context, msg string, err error, fields ...zap.Field) {
	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	l.WithContext(ctx).Error(msg, fields...)
}

func (l *Logger) InfoWithCtx(ctx context.Context, msg string, err error, fields ...zap.Field) {
	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	l.WithContext(ctx).Info(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...any) {
	if len(fields) == 1 {
		l.Logger.Error(msg, zap.Any("data", fields[0]))
		return
	}
	l.Logger.Info(msg, zap.Any("data", fields))
}

func (l *Logger) Error(msg string, err error, fields ...any) {
	if err != nil {
		l.Logger.Error(msg, zap.Error(err), zap.Any("data", fields))
		return
	}
	if len(fields) == 1 {
		l.Logger.Error(msg, zap.Any("data", fields[0]))
		return
	}

	l.Logger.Error(msg, zap.Any("data", fields))
}

func (l *Logger) AppError(msg string, err error) {
	l.Logger.Error(msg, zap.Error(err))
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
