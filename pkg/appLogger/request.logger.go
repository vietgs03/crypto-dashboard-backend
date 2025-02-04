package appLogger

import (
	"fmt"

	"crypto-dashboard-backend/pkg/common"

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

func (logger *Logger) ReqClientLog(method, path string, reqCtx *common.ReqContext) {
	logger.Logger.Info(fmt.Sprintf(`Accepted Request [%s] - %s - %s]`, reqCtx.CID, method, path), zap.Object("data", ReqClientLogger{
		CID:       reqCtx.CID,
		Method:    method,
		Url:       path,
		UserId:    reqCtx.UserInfo.ID,
		TimeStamp: reqCtx.RequestTimestamp,
	}))
}

func (logger *Logger) ResClientLog(reqCtx *common.ReqContext, statusCode int, err error) {
	duration := common.FormatMilliseconds(reqCtx.RequestTimestamp)

	data := ResClientLogger{
		CID:        reqCtx.CID,
		Duration:   duration,
		UserId:     reqCtx.UserInfo.ID,
		StatusCode: statusCode,
	}

	if err != nil {
		data.Error = err.Error()
		logger.Logger.Error(fmt.Sprintf(`Rejected Request [%s] - %d - %s]`, reqCtx.CID, statusCode, duration), zap.Object("data", data))
	} else {
		logger.Logger.Info(fmt.Sprintf(`Response [%s] - %d - %s]`, reqCtx.CID, statusCode, duration), zap.Object("data", data))
	}
}

// TODO : implement server logger for Grpc
// func (u Logger) ReqServerLogger(logger *Logger) {
// 	logger.Logger.Info("", zap.Object("data", u))
// }

// func (u Logger) ResServerLogger(logger *Logger) {
// 	logger.Logger.Info("", zap.Object("data", u))
// }
