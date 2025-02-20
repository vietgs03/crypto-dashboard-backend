package middlewares

import (
	"reflect"

	"crypto-dashboard/gw-example/infrastructure/helper"
	"crypto-dashboard/pkg/common"
	"crypto-dashboard/pkg/constants"
	"crypto-dashboard/pkg/response"

	"github.com/gofiber/fiber/v3"
)

// TODO if Error status 500 to send notfi discord
func ErrorHandler(c fiber.Ctx, err error) error {
	if err != nil && !reflect.ValueOf(err).IsNil() {
		internalError := response.NewAppError(constants.InternalServerErr, err.Error())
		if appErr, ok := err.(*response.AppError); ok {
			internalError = appErr
		}

		c.Status(constants.HttpCode[internalError.Code]).JSON(common.ErrorResponse(helper.GetHttpReqCtx(c), internalError))
	}

	return nil
}
