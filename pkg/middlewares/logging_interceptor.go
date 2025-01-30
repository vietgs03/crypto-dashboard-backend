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

		appLogger.ReqClientLogger{
			CID:       requestContext.CID,
			Method:    c.Method(),
			Url:       c.Path(),
			UserId:    requestContext.UserInfo.ID,
			TimeStamp: requestContext.RequestTimestamp,
		}.Logger(c, logger)
		err := c.Next()
		duration := common.FormatMilliseconds(requestContext.RequestTimestamp)

		appLogger.ResClientLogger{
			CID:        requestContext.CID,
			Duration:   duration,
			StatusCode: c.Response().StatusCode(),
			UserId:     requestContext.UserInfo.ID,
		}.Logger(c, logger, err)

		return err
	}
}
