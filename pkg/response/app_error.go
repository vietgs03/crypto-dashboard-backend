package response

import (
	"fmt"
	"net/http"

	"crypto-dashboard/pkg/constants"
)

type AppError struct {
	Status  int                    `json:"status"`
	Code    constants.InternalCode `json:"code"`
	Data    string                 `json:"data,omitempty"`
	Message string                 `json:"message"`
}

func (r *AppError) IsZero() bool {
	return r.Status == 0
}

func (r *AppError) Error() string {
	return fmt.Sprintf("%d:%s", r.Code, r.Message)
}

func NewAppError(code constants.InternalCode, status int, message string) *AppError {
	return &AppError{Code: code, Message: message, Status: status}
}

func (r *AppError) WithData(data string) *AppError {
	r.Data = data
	return r
}

func QueryNotFound(message string) *AppError {
	return NewAppError(constants.QueryNotFound, http.StatusNotFound, message)
}

func QueryInvalid(message string) *AppError {
	return NewAppError(constants.ParamInvalid, http.StatusInternalServerError, message)
}

func DatabaseError(err error) *AppError {
	return NewAppError(constants.DatabaseErr, http.StatusInternalServerError, err.Error())
}

func MQUnauthorization() *AppError {
	return NewAppError(constants.InvalidToken, http.StatusUnauthorized, "need authenticated for request.")
}

func MQAccessDenined() *AppError {
	return NewAppError(constants.InvalidToken, http.StatusForbidden, "need specific roles for request")
}
