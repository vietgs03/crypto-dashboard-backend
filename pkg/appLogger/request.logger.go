package appLogger

import (
	"fmt"

	"crypto-dashboard/pkg/common"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	ReqClientLogger struct {
		CID       string
		Method    string
		Url       string
		IP        string
		UserId    uint
		TimeStamp int64
	}

	ResClientLogger struct {
		CID        string
		Duration   string
		Error      string
		UserId     uint
		StatusCode uint
		Code       uint
	}
)

func (u ReqClientLogger) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("CID", u.CID)
	enc.AddString("method", u.Method)
	enc.AddString("url", u.Url)
	enc.AddString("ip", u.IP)
	enc.AddUint("userId", u.UserId)
	enc.AddInt64("timeStamp", u.TimeStamp)
	return nil
}

func (u ResClientLogger) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("CID", u.CID)
	enc.AddString("duration", u.Duration)
	enc.AddString("error", u.Error)
	enc.AddUint("userId", u.UserId)
	enc.AddUint("statusCode", u.StatusCode)
	return nil
}

func (logger *Logger) ReqClientLog(reqCtx *common.ReqContext, method, path string) {
	logger.Logger.Info(fmt.Sprintf(`Accepted Request [%s/%s] - %s]`, method, path, reqCtx.CID), zap.Inline(ReqClientLogger{
		CID:       reqCtx.CID,
		Method:    method,
		Url:       path,
		IP:        reqCtx.IP,
		UserId:    reqCtx.UserInfo.ID,
		TimeStamp: reqCtx.RequestTimestamp,
	}))
}

func (logger *Logger) ResClientLog(reqCtx *common.ReqContext, statusCode uint, err error) {
	duration := common.FormatMilliseconds(reqCtx.RequestTimestamp)

	data := ResClientLogger{
		CID:        reqCtx.CID,
		Duration:   duration,
		UserId:     reqCtx.UserInfo.ID,
		StatusCode: statusCode,
		Code:       statusCode,
	}

	if err != nil {
		data.Error = err.Error()
		logger.Logger.Error(fmt.Sprintf(`Rejected Request [%s] - %d - %s]`, reqCtx.CID, statusCode, duration), zap.Inline(data))
	} else {
		logger.Logger.Info(fmt.Sprintf(`Response [%s] - %d - %s]`, reqCtx.CID, statusCode, duration), zap.Inline(data))
	}
}
