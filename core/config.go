package core

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"os"
	"runtime"
	"strconv"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

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
	grpcPort       string
	serverPort     string
	version        bool
)

func init() {
	flag.StringVar(&host, "host", "localhost", "数据库地址")
	flag.StringVar(&user, "user", "root", "数据库用户名")
	flag.StringVar(&password, "password", "123456", "数据库密码")
	flag.StringVar(&port, "port", "3306", "数据库端口")
	flag.StringVar(&database, "database", "trojan_panel_db", "数据库名称")
	flag.StringVar(&accountTable, "accountTable", "account", "account表名称")
	flag.StringVar(&redisHost, "redisHost", "127.0.0.1", "Redis地址")
	flag.StringVar(&redisPort, "redisPort", "6379", "Redis端口")
	flag.StringVar(&redisPassword, "redisPassword", "123456", "Redis密码")
	flag.StringVar(&redisDb, "redisDb", "0", "Redis默认数据库")
	flag.StringVar(&redisMaxIdle, "redisMaxIdle", strconv.FormatInt(int64(runtime.NumCPU()*2), 10), "Redis最大空闲连接数")
	flag.StringVar(&redisMaxActive, "redisMaxActive", strconv.FormatInt(int64(runtime.NumCPU()*2+2), 10), "Redis最大连接数")
	flag.StringVar(&redisWait, "redisWait", "true", "Redis是否等待")
	flag.StringVar(&crtPath, "crtPath", "", "crt秘钥")
	flag.StringVar(&keyPath, "keyPath", "", "key秘钥")
	flag.StringVar(&grpcPort, "grpcPort", "8100", "gRPC端口")
	flag.StringVar(&serverPort, "serverPort", "8082", "服务端口")
	flag.BoolVar(&version, "version", false, "打印版本信息")
	flag.Usage = usage
	flag.Parse()
	if version {
		_, _ = fmt.Fprint(os.Stdout, constant.TrojanPanelCoreVersion)
		os.Exit(0)
	}

	// 初始化日志
	logPath := constant.LogPath
	if !util.Exists(logPath) {
		if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
			logrus.Errorf("创建logs文件夹异常 err: %v", err)
			panic(err)
		}
	}

	// 初始化全局配文件夹
	configPath := constant.ConfigPath
	if !util.Exists(configPath) {
		if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
			logrus.Errorf("创建config文件夹异常 err: %v", err)
			panic(err)
		}
	}

	configFilePath := constant.ConfigFilePath
	if !util.Exists(configFilePath) {
		file, err := os.Create(configFilePath)
		if err != nil {
			logrus.Errorf("创建config.ini文件异常 err: %v", err)
			panic(err)
		}
		defer file.Close()

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
[grpc]
port=%s
[server]
port=%s
`, host, user, password, port, database, accountTable, redisHost, redisPort, redisPassword, redisDb,
			redisMaxIdle, redisMaxActive, redisWait, crtPath, keyPath, grpcPort, serverPort))
		if err != nil {
			logrus.Errorf("config.ini文件写入异常 err: %v", err)
			panic(err)
		}
	}

	sqlitePath := constant.SqlitePath
	if !util.Exists(sqlitePath) {
		if err := os.MkdirAll(sqlitePath, os.ModePerm); err != nil {
			logrus.Errorf("创建sqlite文件夹异常 err: %v", err)
			panic(err)
		}
	}
	sqliteFilePath := constant.SqliteFilePath
	if !util.Exists(sqliteFilePath) {
		file, err := os.Create(sqliteFilePath)
		if err != nil {
			logrus.Errorf("创建node_config.db文件异常 err: %v", err)
			panic(err)
		}
		defer file.Close()
	}
}

func usage() {
	_, _ = fmt.Fprintln(os.Stdout, `trojan panel core manage help
Usage: trojan-panel-core [-host] [-user] [-password] [-port] [-database] [-accountTable] [-redisHost] [-redisPort] [-redisPassword] [-redisDb] [-redisMaxIdle] [-redisMaxActive] [-redisWait] [-crtPath] [-keyPath] [-grpcPort] [-serverPort] [-h] [-version]`)
	flag.PrintDefaults()
}

var Config = new(AppConfig)

// InitConfig 初始化全局配置文件
func InitConfig() {
	if err := ini.MapTo(Config, constant.ConfigFilePath); err != nil {
		logrus.Errorf("配置文件加载失败 err: %v", err)
		panic(err)
	}
}

type AppConfig struct {
	MySQLConfig  `ini:"mysql"`
	RedisConfig  `ini:"redis"`
	CertConfig   `ini:"cert"`
	LogConfig    `ini:"log"`
	GrpcConfig   `ini:"grpc"`
	ServerConfig `ini:"server"`
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

// GrpcConfig gRPC
type GrpcConfig struct {
	Port string `ini:"port"` // gRPC端口
}

type ServerConfig struct {
	Port int `ini:"port"` // 服务器端口
}
