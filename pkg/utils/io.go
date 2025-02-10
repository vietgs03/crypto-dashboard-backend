package utils

import (
	"os"
	"path/filepath"
)

func GetFileName(filePath string) (string, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}
	return fileInfo.Name(), nil
}

func GetDirectoryPath(filePath string) string {
	dirPath := filepath.Dir(filePath)
	return dirPath
}
