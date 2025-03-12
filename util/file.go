package util

import (
	"fmt"
	"io"
	"os"
)

func RemoveFile(filePath string) error {
	if Exists(filePath) {
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("failed to delete file")
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

func SaveBytesToFile(data []byte, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if _, err = io.WriteString(file, string(data)); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}
