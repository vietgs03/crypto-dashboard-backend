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
	_ = os.WriteFile(name, b, 0644)
}
