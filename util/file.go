package util

import (
	"errors"
	"os"
)

func RemoveFile(filePath string) error {
	if Exists(filePath) {
		if err := os.Remove(filePath); err != nil {
			return errors.New("failed to delete file")
		}
	}
	return nil
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
