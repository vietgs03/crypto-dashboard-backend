package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func CalcSHA256Digest(content string) string {
	hash := sha256.New()
	hash.Write([]byte(content))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func EncodeWithBase64(value []byte) string {
	return base64.URLEncoding.EncodeToString(value)
}

func MD5Hash(content string) string {
	data := []byte(content)
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)
}
