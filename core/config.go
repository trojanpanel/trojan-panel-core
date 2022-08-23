package core

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"trojan-panel-core/module/constant"
)

var Config = new(AppConfig)

// InitConfig 初始化全局配置文件
func InitConfig() {
	if err := ini.MapTo(Config, constant.ConfigFilePath); err != nil {
		logrus.Errorf("配置文件加载失败 err: %v\n", err)
		panic(err)
	}
}

type AppConfig struct {
	MySQLConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
	CertConfig  `ini:"cert"`
	LogConfig   `ini:"log"`
}

// MySQLConfig MySQL
type MySQLConfig struct {
	Host         string `ini:"host"`
	User         string `ini:"user"`
	Password     string `ini:"password"`
	Port         int    `ini:"port"`
	Database     string `ini:"database"`
	AccountTable string `ini:"account_table"`
}

type RedisConfig struct {
	Host      string `ini:"host"`
	Port      int    `ini:"port"`
	Password  string `ini:"password"`
	Db        int    `ini:"db"`
	MaxIdle   int    `ini:"max_idle"`
	MaxActive int    `ini:"max_active"`
	Wait      bool   `ini:"wait"`
}

type CertConfig struct {
	CrtPath string `ini:"crt_path"`
	KeyPath string `ini:"key_path"`
}

// LogConfig log
type LogConfig struct {
	FileName   string `ini:"filename"`    // 日志文件位置
	MaxSize    int    `ini:"max_size"`    // 单文件最大容量,单位是MB
	MaxBackups int    `ini:"max_backups"` // 最大保留过期文件个数
	MaxAge     int    `ini:"max_age"`     // 保留过期文件的最大时间间隔,单位是天
	Compress   bool   `ini:"compress"`    // 是否需要压缩滚动日志, 使用的 gzip 压缩
}
