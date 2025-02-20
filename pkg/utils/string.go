package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"crypto-dashboard/pkg/constants"
	"crypto-dashboard/pkg/response"

	"github.com/samber/lo"
	"github.com/sqids/sqids-go"
)

var sqID *sqids.Sqids

func init() {
	if sqID == nil {
		sqID, _ = sqids.New(sqids.Options{
			Alphabet:  "FxnXM1kBN6cuhsAvjW3Co7l2RePyY8DwaU04Tzt9fHQrqSVKdpimLGIJOgb5ZE",
			MinLength: 8,
		})
	}
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomNumber(length int) string {
	const charset = "0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateShortUUID(baseId uint) string {
	id, _ := sqID.Encode([]uint64{uint64(baseId)})
	return id
}

func ParseIdFromShortUUID(uuid string) (*uint, *response.AppError) {
	ids := sqID.Decode(uuid)
	if len(ids) == 0 {
		return nil, response.NewAppError(constants.UuidInvalid, "Invalid short UUID")
	}

	return lo.ToPtr(uint(ids[0])), nil
}

// Example:  `[{"name": "name1"}, {"name": "name2"}, {"name": "name3"}]`
func StringToArrayStruct[T any](s string) ([]T, error) {
	var t []T

	err := json.Unmarshal([]byte(s), &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// Example: {"names": ["name1"], "explanation_1": "name1", "explanation_2": "name1"}
func StringToStruct[T any](s string) (*T, error) {
	var t T

	err := json.Unmarshal([]byte(s), &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func StructToString[T any](s T) (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func JoinComma(s []string) string {
	if len(s) == 0 {
		return ""
	}

	return strings.Join(s, ",")
}

type convertible interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64
}

func ToString[T convertible](c T) string {
	return fmt.Sprintf("%v", c)
}

// SliceJoinComma convert []uint to a comma-separated string
func SliceJoinComma[T convertible](s []T, separated string) string {
	strIds := make([]string, 0, len(s))
	for _, v := range s {
		strIds = append(strIds, fmt.Sprintf("%v", v))
	}
	return strings.Join(strIds, separated)
}
