package middlewares

import (
	"crypto-dashboard/pkg/constants"
	"crypto-dashboard/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// TODO if Error status 500 to send notfi discord
func ErrorHandler(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		if appErr, ok := err.(*response.AppError); ok {
			response.ErrorResponse(c, appErr)
		}
		internalError := response.NewAppError(constants.InternalServerErr, fiber.StatusInternalServerError, err.Error())
		response.ErrorResponse(c, internalError)
	}

	return nil
}
