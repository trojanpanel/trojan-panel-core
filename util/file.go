package util

import (
	"fmt"
	"io"
	"net/http"
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

func DownloadFromGithub(binName, binPath, owner, repo, version string) error {
	url, err := GetReleaseAssetURL(owner, repo, version, binName)
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file, status code: %d", resp.StatusCode)
	}

	if Exists(binPath) {
		if err = os.Remove(binPath); err != nil {
			return fmt.Errorf("failed to remove existing file: %v", err)
		}
	}

	file, err := os.Create(binPath)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", binPath, err)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	if err = os.Chmod(binPath, 0755); err != nil {
		return fmt.Errorf("failed to change file permissions: %v", err)
	}
	return nil
}
