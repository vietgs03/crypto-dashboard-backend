package common

import (
	"fmt"
	"time"

	"crypto-dashboard/pkg/dtos/duser"

	"github.com/google/uuid"
)

type (
	ReqContext struct {
		CID              string `json:"cid"`
		IP               string
		RequestTimestamp int64 `json:"request_timestamp"`
		UserInfo         *duser.UserInfo
		AccessToken      *string `json:"access_token"`
		RefreshToken     *string `json:"refresh_token"`
	}
)

func BuildRequestContext(cid *string, accessToken *string, refreshToken *string, userInfo *duser.UserInfo) *ReqContext {
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
