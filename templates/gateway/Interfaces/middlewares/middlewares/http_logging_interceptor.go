package middlewares

import (
	"crypto-dashboard/pkg/constants"

	"crypto-dashboard/pkg/appLogger"
	"crypto-dashboard/pkg/common"

	"github.com/gofiber/fiber/v3"
)

func LoggingInterceptor(logger *appLogger.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		requestContext := c.Locals(constants.REQUEST_CONTEXT_KEY).(*common.ReqContext)

		logger.ReqClientLog(requestContext, c.Method(), c.Path())

		err := c.Next()

		logger.ResClientLog(requestContext, uint(c.Response().StatusCode()), err)
		return err
	}
}
