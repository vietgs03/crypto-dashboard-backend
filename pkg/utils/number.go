package utils

import (
	"strconv"
)

func ToInt[T int | int8 | int16 | int32 | int64](s string) T {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}

	return T(val)
}
