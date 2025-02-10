package middlewares

import (
	"crypto-dashboard-backend/pkg/appLogger"
	"crypto-dashboard-backend/pkg/common"
	"crypto-dashboard-backend/pkg/constants"

	"github.com/gofiber/fiber/v2"
)

func LoggingInterceptor(logger *appLogger.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestContext := c.Locals(constants.REQUEST_CONTEXT_KEY).(*common.ReqContext)

		logger.ReqClientLog(c.Method(), c.Path(), requestContext)

		err := c.Next()

		logger.ResClientLog(requestContext, c.Response().StatusCode(), err)
		return err
	}
}
