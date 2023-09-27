package util

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"trojan-panel-core/model/constant"
)

func DownloadFile(url string, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = os.WriteFile(fileName, data, 0755); err != nil {
		return err
	}
	return nil
}

func RemoveFile(fileName string) error {
	if Exists(fileName) {
		if err := os.Remove(fileName); err != nil {
			logrus.Errorf("failed to delete file fileName: %s err: %v", fileName, err)
			return errors.New(constant.RemoveFileError)
		}
	}
	return nil
}

// Unzip decompress
func Unzip(src string, dest string) error {
	// open and read compressed files
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if r != nil {
			r.Close()
		}
	}()

	// traverse the files in the compressed file and write to disk
	for _, f := range r.File {
		filePath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(filePath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: invalid file path", filePath)
		}

		// If it is a directory, create the directory
		if f.FileInfo().IsDir() {
			if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}

		rc.Close()
		outFile.Close()
	}
	return nil
}

// Exists check if a file or folder exists
func Exists(path string) bool {
	// get file info
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
