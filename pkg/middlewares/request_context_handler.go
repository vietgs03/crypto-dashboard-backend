package middlewares

import (
	"crypto-dashboard-backend/pkg/common"
	"crypto-dashboard-backend/pkg/constants"

	"github.com/gofiber/fiber/v2"
)

func ReqContextHandler(c *fiber.Ctx) error {
	cid := c.Get(string(constants.CORRELATION_ID_KEY))

	authorization := c.Get(string(constants.AUTHORIZATION_KEY))
	if authorization != "" {
		authorization = authorization[7:]
	}

	requestContext := common.BuildRequestContext(&cid, &authorization, nil, nil)

	c.Locals(constants.REQUEST_CONTEXT_KEY, requestContext)
	c.Locals(constants.CORRELATION_ID_KEY, requestContext.CID)

	return c.Next()
}
