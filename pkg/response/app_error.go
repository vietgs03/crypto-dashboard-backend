package response

import (
	"fmt"
	"reflect"

	"crypto-dashboard/pkg/constants"
)

type AppError struct {
	Code    constants.InternalCode `json:"code"`
	Data    string                 `json:"data,omitempty"`
	Message string                 `json:"message"`
}

func (r *AppError) IsZero() bool {
	return r.Code == 0
}

func (r *AppError) Error() string {
	return fmt.Sprintf("%d:%s", r.Code, r.Message)
}

func NewAppError(code constants.InternalCode, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func (r *AppError) WithData(data string) *AppError {
	r.Data = data
	return r
}

func QueryNotFound(message string) *AppError {
	return NewAppError(constants.QueryNotFound, message)
}

func QueryInvalid(message string) *AppError {
	if message == "" {
		message = "invalid query"
	}
	return NewAppError(constants.ParamInvalid, message)
}

func DatabaseError(err error) *AppError {
	return NewAppError(constants.DatabaseErr, err.Error())
}

func Unauthorization(message string) *AppError {
	if message != "" {
		message = "need authenticated for request."
	}
	return NewAppError(constants.InvalidToken, "need authenticated for request.")
}

func AccessDenined() *AppError {
	return NewAppError(constants.AcceptDenied, "need specific roles for request")
}

func UnknownError(message string) *AppError {
	if message == "" {
		message = "unknown error"
	}
	return NewAppError(constants.InternalServerErr, message)
}

func ServerError(message string) *AppError {
	if message == "" {
		message = "server error"
	}
	return NewAppError(constants.InternalServerErr, message)
}

func ConvertError(err error) *AppError {
	if err == nil || reflect.TypeOf(err).Kind() == reflect.Ptr && reflect.ValueOf(err).IsNil() {
		return nil
	}

	if appErr, ok := err.(*AppError); ok {
		return appErr
	}

	return NewAppError(constants.InternalServerErr, err.Error())
}
