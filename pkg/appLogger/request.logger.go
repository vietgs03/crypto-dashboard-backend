package appLogger

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	ReqClientLogger struct {
		CID       string
		Method    string
		Url       string
		UserId    uint
		TimeStamp int64
	}

	ResClientLogger struct {
		CID        string
		Duration   string
		Error      string
		UserId     uint
		StatusCode int
		Code       int
	}

	ReqServerLogger struct {
		EventName string
		CID       string
		UserId    uint
		TimeStamp int64
	}

	ResServerLogger struct {
		CID        string
		UserId     uint
		Error      string
		StatusCode int
		Duration   string
	}
)

func (u ReqClientLogger) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("CID", u.CID)
	enc.AddString("method", u.Method)
	enc.AddString("url", u.Url)
	enc.AddUint("userId", u.UserId)
	enc.AddInt64("timeStamp", u.TimeStamp)
	return nil
}

func (u ResClientLogger) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("CID", u.CID)
	enc.AddString("duration", u.Duration)
	enc.AddString("error", u.Error)
	enc.AddUint("userId", u.UserId)
	enc.AddInt("statusCode", u.StatusCode)
	return nil
}

func (u ReqServerLogger) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("CID", u.CID)
	enc.AddString("eventName", u.EventName)
	enc.AddUint("userId", u.UserId)
	enc.AddInt64("timeStamp", u.TimeStamp)
	return nil
}

func (u ResServerLogger) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("CID", u.CID)
	enc.AddString("duration", u.Duration)
	enc.AddString("error", u.Error)
	enc.AddUint("userId", u.UserId)
	enc.AddInt("statusCode", u.StatusCode)
	return nil
}

func (u ReqClientLogger) Logger(c *fiber.Ctx, logger *Logger) {
	logger.Logger.Info(fmt.Sprintf(`Accepted Request [%s] - %s - %s]`, u.CID, u.Method, u.Url), zap.Object("data", u))
}

func (u ResClientLogger) Logger(c *fiber.Ctx, logger *Logger, err error) {
	// TODO how get internal code
	if err != nil {
		u.Error = err.Error()
		logger.Logger.Info(fmt.Sprintf(`Rejected Request [%s] - %d - %s]`, u.CID, u.StatusCode, u.Duration), zap.Object("data", u))
	} else {
		logger.Logger.Error(fmt.Sprintf(`Response [%s] - %d - %s]`, u.CID, u.StatusCode, u.Duration), zap.Object("data", u))
	}
}

// TODO : implement server logger for Grpc
func (u ReqServerLogger) Logger(logger *Logger) {
	logger.Logger.Info("", zap.Object("data", u))
}

func (u ResServerLogger) Logger(logger *Logger) {
	logger.Logger.Info("", zap.Object("data", u))
}
