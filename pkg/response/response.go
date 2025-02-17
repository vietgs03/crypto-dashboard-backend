package response

import (
	"crypto-dashboard/pkg/constants"

	"github.com/gofiber/fiber/v2"
)

type ResponseData struct {
	Cid     string      `json:"cid"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponseData struct {
	Cid    string      `json:"cid"`
	Code   int         `json:"code"`
	Err    string      `json:"error"`
	Detail interface{} `json:"detail"`
}

// success response
func SuccessResponse(c *fiber.Ctx, statusCode int, code int, data interface{}) {
	c.Status(statusCode).JSON(ResponseData{
		Cid:     c.Locals(constants.CORRELATION_ID_KEY).(string),
		Code:    int(constants.Success),
		Message: "success",
		Data:    data,
	})
}

func ErrorResponse(c *fiber.Ctx, err *AppError) {
	var message string
	if err.Message == "" {
		message = constants.Msg[err.Code]
	} else {
		message = constants.Msg[constants.InternalServerErr]
	}
	c.Status(err.Status).JSON(ErrorResponseData{
		Cid:    c.Locals(constants.CORRELATION_ID_KEY).(string),
		Code:   int(err.Code),
		Err:    message,
		Detail: err.Data,
	})
}
