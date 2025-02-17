package appLogger

import (
	"fmt"

	"crypto-dashboard/pkg/common"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	ReqServiceLogger struct {
		CID        string
		Method     string
		Url        string
		TimeStamp  int64
		ServerName string
	}

	ResServiceLogger struct {
		CID        string
		Duration   string
		Error      string
		ServerName string
		StatusCode uint
		Code       uint
	}
)

func (u ReqServiceLogger) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("CID", u.CID)
	enc.AddString("serverName", u.ServerName)
	enc.AddString("method", u.Method)
	enc.AddString("url", u.Url)
	enc.AddInt64("timeStamp", u.TimeStamp)
	return nil
}

func (u ResServiceLogger) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("CID", u.CID)
	enc.AddString("duration", u.Duration)
	enc.AddString("error", u.Error)
	enc.AddUint("statusCode", u.StatusCode)
	return nil
}

// TODO : implement server logger for Grpc
func (u *Logger) ReqServerLogger(reqCtx *common.ReqContext, method, path string) {
	u.Logger.Info(fmt.Sprintf(`Accepted Request [%s/%s] - %s]`, method, path, reqCtx.CID), zap.Inline(ReqServiceLogger{
		ServerName: u.serviceName,
		Method:     method,
		Url:        path,
		CID:        reqCtx.CID,
		TimeStamp:  reqCtx.RequestTimestamp,
	}))
}

func (u *Logger) ResServerLogger(reqCtx *common.ReqContext, statusCode uint, err error) {
	duration := common.FormatMilliseconds(reqCtx.RequestTimestamp)

	data := ResServiceLogger{
		CID:        reqCtx.CID,
		Duration:   duration,
		ServerName: u.serviceName,
		StatusCode: statusCode,
	}
	if err != nil {
		data.Error = err.Error()
		u.Logger.Error(fmt.Sprintf(`Rejected Request [%s] - %d - %s]`, reqCtx.CID, statusCode, duration), zap.Inline(data))
	} else {
		u.Logger.Info(fmt.Sprintf(`Response [%s] - %d - %s]`, reqCtx.CID, statusCode, duration), zap.Inline(data))
	}
}
