package util

import (
	"archive/zip"
	"errors"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"trojan-panel-core/module/constant"
)

var configFileNameReg = regexp.MustCompile("^config-([1-9]\\d*)\\.json$")

func InitFile() {
	// 初始化日志
	logPath := constant.LogPath
	if !Exists(logPath) {
		if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
			logrus.Errorf("创建logs文件夹异常 err: %v\n", err)
			panic(err)
		}
	}

	// 初始化bin文件夹
	InitBinFilePath()

	// 初始化全局配置文件
	InitConfigFile()
}

func InitConfigFile() {
	// 初始化全局配文件夹
	configPath := constant.ConfigPath
	if !Exists(configPath) {
		if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
			logrus.Errorf("创建config文件夹异常 err: %v\n", err)
			panic(err)
		}
	}

	configFilePath := constant.ConfigFilePath
	if !Exists(configFilePath) {
		file, err := os.Create(configFilePath)
		if err != nil {
			logrus.Errorf("创建config.ini文件异常 err: %v\n", err)
			panic(err)
		}
		defer file.Close()

		var (
			host           string
			user           string
			password       string
			port           string
			database       string
			accountTable   string
			redisHost      string
			redisPort      string
			redisPassword  string
			redisDb        string
			redisMaxIdle   string
			redisMaxActive string
			redisWait      string
			crtPath        string
			keyPath        string
		)
		flag.StringVar(&host, "host", "localhost", "数据库地址")
		flag.StringVar(&user, "user", "root", "数据库用户名")
		flag.StringVar(&password, "password", "123456", "数据库密码")
		flag.StringVar(&port, "port", "3306", "数据库端口")
		flag.StringVar(&database, "database", "trojan_panel_db", "数据库名称")
		flag.StringVar(&accountTable, "account-table", "account", "account表名称")
		flag.StringVar(&redisHost, "redisHost", "127.0.0.1", "Redis地址")
		flag.StringVar(&redisPort, "redisPort", "6379", "Redis端口")
		flag.StringVar(&redisPassword, "redisPassword", "123456", "Redis密码")
		flag.StringVar(&redisDb, "redisDb", "0", "Redis默认数据库")
		flag.StringVar(&redisMaxIdle, "redisMaxIdle", "2", "Redis最大空闲连接数")
		flag.StringVar(&redisMaxActive, "redisMaxActive", "2", "Redis最大连接数")
		flag.StringVar(&redisWait, "redisWait", "true", "Redis是否等待")
		flag.StringVar(&crtPath, "crt-path", "", "crt秘钥")
		flag.StringVar(&keyPath, "key-path", "", "key秘钥")
		flag.Parse()
		_, err = file.WriteString(fmt.Sprintf(
			`[mysql]
host=%s
user=%s
password=%s
port=%s
database=%s
account_table=%s
[redis]
host=%s
port=%s
password=%s
db=%s
max_idle=%s
max_active=%s
wait=%s
[cert]
crt_path=%s
key_path=%s
[log]
filename=logs/trojan-panel-core.log
max_size=1
max_backups=5
max_age=30
compress=true
`, host, user, password, port, database, accountTable, redisHost, redisPort, redisPassword, redisDb,
			redisMaxIdle, redisMaxIdle, redisWait, crtPath, keyPath))
		if err != nil {
			logrus.Errorf("config.ini文件写入异常 err: %v\n", err)
			panic(err)
		}
		flag.Usage = usage
	}
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `trojan panel core manage help
Usage: trojan-panel-core [-host] [-password] [-port] [-database] [-account-table] [-redisHost] [-redisPort] [-redisPassword] [-redisDb] [-redisMaxIdle] [-redisMaxActive] [-redisWait] [-crt-path] [-key-path] [-h]

Options:
-host            database host
-user            database user
-password        database password
-port            database port
-database        database name
-account-table   account table name
-redisHost		 redis redisHost
-redisPort		 redis redisPort
-redisPassword	 redis redisPassword
-redisDb		 redis redisDb
-redisMaxIdle    redis redisMaxIdle
-redisMaxActive	 redis redisMaxActive
-redisWait		 redis redisWait
-crt-path	 	 cert crt file path
-key-path	 	 cert key file path
-h               help
`)
}

func DownloadFile(url string, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

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
	if err := os.Remove(fileName); err != nil {
		logrus.Errorf("删除文件失败 fileName: %s err: %v\n", fileName, err)
		return errors.New(constant.RemoveFileError)
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
	defer r.Close()

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

func GetConfigApiPorts(dirPth string) ([]uint, error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	apiPorts := make([]uint, 0)
	for _, fi := range dir {
		// 过滤指定格式
		finds := configFileNameReg.FindStringSubmatch(fi.Name())
		if len(finds) > 0 {
			apiPort, err := strconv.Atoi(finds[1])
			if err != nil {
				logrus.Errorf("类型转换异常 err: %v\n", err)
				continue
			}
			apiPorts = append(apiPorts, uint(apiPort))
		}
	}
	return apiPorts, nil
}

func InitBinFilePath() {
	xrayPath := constant.XrayPath
	if !Exists(xrayPath) {
		if err := os.MkdirAll(xrayPath, os.ModePerm); err != nil {
			logrus.Errorf("创建/bin/xray文件夹异常 err: %v\n", err)
			panic(err)
		}
	}
	trojanGoPath := constant.TrojanGoPath
	if !Exists(trojanGoPath) {
		if err := os.MkdirAll(trojanGoPath, os.ModePerm); err != nil {
			logrus.Errorf("创建/bin/trojango文件夹异常 err: %v\n", err)
			panic(err)
		}
	}
	hysteriaPath := constant.HysteriaPath
	if !Exists(hysteriaPath) {
		if err := os.MkdirAll(hysteriaPath, os.ModePerm); err != nil {
			logrus.Errorf("创建/bin/hysteriaPath文件夹异常 err: %v\n", err)
			panic(err)
		}
	}
}
