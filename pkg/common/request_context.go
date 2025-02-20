package common

import (
	"context"
	"fmt"
	"time"

	"crypto-dashboard/pkg/constants"
	"crypto-dashboard/pkg/dtos/duser"
	"crypto-dashboard/pkg/utils"

	"github.com/google/uuid"
)

type (
	ReqContext struct {
		CID              string `json:"cid"`
		IP               string
		RequestTimestamp int64 `json:"request_timestamp"`
		UserInfo         *duser.UserInfo[any]
		AccessToken      *string `json:"access_token"`
		RefreshToken     *string `json:"refresh_token"`
	}
)

func BuildRequestContext(cid *string, accessToken *string, refreshToken *string, userInfo *duser.UserInfo[any]) *ReqContext {
	return &ReqContext{
		CID:              GetCid(cid),
		RequestTimestamp: time.Now().UnixMilli(),
		AccessToken:      accessToken,
		UserInfo:         userInfo,
		RefreshToken:     refreshToken,
	}
}

func GetCid(cid *string) string {
	if cid == nil {
		return ""
	}

	newCid, err := uuid.NewV7()
	if err != nil {
		return uuid.NewString()
	}
	return newCid.String()
}

func CalculateDuration(timeStamp int64) int64 {
	return time.Now().UnixMilli() - timeStamp
}

func FormatMilliseconds(timeStamp int64) string {
	return fmt.Sprintf("%dms", CalculateDuration(timeStamp))
}

func GetUserCtx[T any](ctx context.Context) *duser.UserInfo[T] {
	reqCtx := GetReqCtx(ctx)
	res, _ := utils.StructToStruct[duser.UserInfo[any], duser.UserInfo[T]](reqCtx.UserInfo)

	return res
}

func GetReqCtx(ctx context.Context) *ReqContext {
	reqCtx := ctx.Value(constants.REQUEST_CONTEXT_KEY).(*ReqContext)

	if reqCtx == nil {
		return BuildRequestContext(nil, nil, nil, &duser.UserInfo[any]{})
	}

	return reqCtx
}

func SetUserCtx(ctx context.Context, user *duser.UserInfo[any]) context.Context {
	reqCtx := ctx.Value(constants.REQUEST_CONTEXT_KEY).(*ReqContext)

	if reqCtx == nil {
		reqCtx = BuildRequestContext(nil, nil, nil, user)
		ctx = context.WithValue(ctx, constants.REQUEST_CONTEXT_KEY, reqCtx)
	} else {
		reqCtx.UserInfo = user
	}

	return ctx
}
