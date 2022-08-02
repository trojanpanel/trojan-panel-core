package util

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"xray-manage/module/constant"
)

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

func InitFile() {
	logPath := constant.LogPath
	if !Exists(logPath) {
		if err := os.Mkdir(logPath, os.ModePerm); err != nil {
			logrus.Errorf("创建logs文件夹异常 err: %v\n", err)
			panic(err)
		}
	}

	xrayConfigPath := constant.XrayConfigPath
	if !Exists(xrayConfigPath) {
		if err := os.Mkdir(xrayConfigPath, os.ModePerm); err != nil {
			logrus.Errorf("创建xray文件夹异常 err: %v\n", err)
			panic(err)
		}
	}

	xrayConfigFilePath := constant.XrayConfigFilePath
	if !Exists(xrayConfigFilePath) {
		file, err := os.Create(xrayConfigFilePath)
		if err != nil {
			logrus.Errorf("创建xray config.json文件异常 err: %v\n", err)
			panic(err)
		}
		defer file.Close()

		_, err = file.WriteString(`{
  "stats": {},
  "api": {
    "services": [
      "HandlerService",
      "LoggerService",
      "StatsService"
    ],
    "tag": "api"
  },
  "policy": {
    "levels": {
      "0": {
        "statsUserUplink": true,
        "statsUserDownlink": true
      }
    },
    "system": {
      "statsInboundUplink": true,
      "statsInboundDownlink": true,
      "statsOutboundUplink": true,
      "statsOutboundDownlink": true
    }
  },
  "inbounds": [
    {
      "tag": "api",
      "listen": "127.0.0.1",
      "port": 10087,
      "protocol": "dokodemo-door",
      "settings": {
        "address": "127.0.0.1"
      }
    }
  ],
  "outbounds": [
    {
      "tag": "direct",
      "protocol": "freedom",
      "settings": {}
    }
  ],
  "routing": {
    "rules": [
      {
        "inboundTag": [
          "api"
        ],
        "outboundTag": "api",
        "type": "field"
      }
    ]
  }
}
`)
		if err != nil {
			logrus.Errorf("xray config.json文件写入异常 err: %v\n", err)
			panic(err)
		}
	}

	configPath := constant.ConfigPath
	if !Exists(configPath) {
		if err := os.Mkdir(configPath, os.ModePerm); err != nil {
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
			host     string
			user     string
			password string
			port     string
			database string
		)
		flag.StringVar(&host, "host", "localhost", "数据库地址")
		flag.StringVar(&user, "user", "root", "数据库用户名")
		flag.StringVar(&password, "password", "123456", "数据库密码")
		flag.StringVar(&port, "port", "3306", "数据库端口")
		flag.StringVar(&database, "database", "trojan_panel_db", "数据库名称")
		flag.Parse()
		_, err = file.WriteString(fmt.Sprintf(
			`[mysql]
		host =%s
		user =%s
		password =%s
		port =%s
		database =%s
[log]
filename = logs/xray-manage.log
max_size = 1
max_backups = 5
max_age = 30
compress = true
`, host, user, password, port, database))
		if err != nil {
			logrus.Errorf("config.ini文件写入异常 err: %v\n", err)
			panic(err)
		}
		flag.Usage = usage
	}
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `xray manage help
Usage: xraymanage [-host] [-password] [-port] [-database]

Options:
-host            database host
-user            database user
-password        database password
-port            database port
-database        database name
-h                help
`)
}
