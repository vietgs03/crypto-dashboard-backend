package common

import (
	"crypto-dashboard/pkg/constants"
	"crypto-dashboard/pkg/response"
)

type ResponseData struct {
	CID     string `json:"cid"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponseData struct {
	CID    string `json:"cid"`
	Code   int    `json:"code"`
	Err    string `json:"error"`
	Detail any    `json:"detail"`
}

// success response
func SuccessResponse(reqCtx *ReqContext, code int, data any) *ResponseData {
	return &ResponseData{
		CID:     reqCtx.CID,
		Code:    int(constants.Success),
		Message: "success",
		Data:    data,
	}
}

func ErrorResponse(reqCtx *ReqContext, err *response.AppError) *ErrorResponseData {
	var message string
	if err.Message == "" {
		message = constants.Msg[err.Code]
	} else {
		message = constants.Msg[constants.InternalServerErr]
	}

	return &ErrorResponseData{
		CID:    reqCtx.CID,
		Code:   int(err.Code),
		Err:    message,
		Detail: err.Data,
	}
}
