package utils

import (
	"encoding/json"
	"os"
)

func IsEmpty[T any](arr []T) bool {
	if arr == nil {
		return true
	}

	return len(arr) == 0
}

func IsBlank(v *string) bool {
	if v == nil {
		return true
	}

	return len(*v) == 0
}

func WriteFile(name string, data any) {
	b, _ := json.Marshal(data)
	_ = os.WriteFile(name, b, 0o644)
}

func PaginationOpts(page, limit uint) (take, skip uint) {
	limit = min(limit, 100)

	page = max(page, 1)
	return limit, (page - 1) * limit
}

func max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}
