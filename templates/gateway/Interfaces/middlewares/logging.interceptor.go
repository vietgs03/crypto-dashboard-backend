package middlewares

import (
	"crypto-dashboard/gw-example/infrastructure/global"
	"crypto-dashboard/pkg/constants"

	"crypto-dashboard/pkg/common"

	"github.com/gofiber/fiber/v3"
)

func LoggingInterceptor(c fiber.Ctx) error {
	requestContext := c.Locals(constants.REQUEST_CONTEXT_KEY).(*common.ReqContext)

	global.Log.ReqClientLog(requestContext, c.Method(), c.Path())

	err := c.Next()

	global.Log.ResClientLog(requestContext, uint(c.Response().StatusCode()), err)
	return err
}
