package util

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"trojan-panel-core/module/constant"
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

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(fileName, data, 0755); err != nil {
		return err
	}
	return nil
}

func RemoveFile(fileName string) error {
	if Exists(fileName) {
		if err := os.Remove(fileName); err != nil {
			logrus.Errorf("删除文件失败 fileName: %s err: %v", fileName, err)
			return errors.New(constant.RemoveFileError)
		}
	}
	return nil
}

// Unzip 解压
func Unzip(src string, dest string) error {
	// 打开读取压缩文件
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if r != nil {
			r.Close()
		}
	}()

	// 遍历压缩文件内的文件，写入磁盘
	for _, f := range r.File {
		filePath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(filePath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: 非法的文件路径", filePath)
		}

		// 如果是目录，就创建目录
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

// Exists 判断文件或者文件夹是否存在
func Exists(path string) bool {
	// 获取文件信息
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
