package helper

import (
	"crypto-dashboard/pkg/common"
	"crypto-dashboard/pkg/constants"
	"crypto-dashboard/pkg/dtos/duser"

	"github.com/gofiber/fiber/v3"
)

func GetHttpUserCtx(c fiber.Ctx) *duser.UserInfo[any] {
	return c.Locals(constants.REQUEST_CONTEXT_KEY).(*common.ReqContext).UserInfo
}

func GetHttpReqCtx(c fiber.Ctx) *common.ReqContext {
	return c.Locals(constants.REQUEST_CONTEXT_KEY).(*common.ReqContext)
}

func SetHttpUserCtx(c fiber.Ctx, user *duser.UserInfo[any]) {
	c.Locals(constants.REQUEST_CONTEXT_KEY).(*common.ReqContext).UserInfo = user
}
