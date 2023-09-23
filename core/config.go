package core

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"os"
	"runtime"
	"strconv"
	"trojan-panel-core/model/constant"
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
	flag.StringVar(&host, "host", "localhost", "database address")
	flag.StringVar(&user, "user", "root", "database username")
	flag.StringVar(&password, "password", "123456", "database password")
	flag.StringVar(&port, "port", "3306", "database port")
	flag.StringVar(&database, "database", "trojan_panel_db", "database name")
	flag.StringVar(&accountTable, "accountTable", "account", "account table name")
	flag.StringVar(&redisHost, "redisHost", "127.0.0.1", "redis address")
	flag.StringVar(&redisPort, "redisPort", "6379", "redis port")
	flag.StringVar(&redisPassword, "redisPassword", "123456", "redis password")
	flag.StringVar(&redisDb, "redisDb", "0", "redis default database")
	flag.StringVar(&redisMaxIdle, "redisMaxIdle", strconv.FormatInt(int64(runtime.NumCPU()*2), 10), "redis maximum number of idle connections")
	flag.StringVar(&redisMaxActive, "redisMaxActive", strconv.FormatInt(int64(runtime.NumCPU()*2+2), 10), "redis maximum number of connections")
	flag.StringVar(&redisWait, "redisWait", "true", "does Redis wait")
	flag.StringVar(&crtPath, "crtPath", "", "crt cert")
	flag.StringVar(&keyPath, "keyPath", "", "key cert")
	flag.StringVar(&grpcPort, "grpcPort", "8100", "gRPC port")
	flag.StringVar(&serverPort, "serverPort", "8082", "service port")
	flag.BoolVar(&version, "version", false, "print version info")
	flag.Usage = usage
	flag.Parse()
	if version {
		_, _ = fmt.Fprint(os.Stdout, constant.TrojanPanelCoreVersion)
		os.Exit(0)
	}

	// initialization log
	logPath := constant.LogPath
	if !util.Exists(logPath) {
		if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
			logrus.Errorf("create logs folder err: %v", err)
			panic(err)
		}
	}

	// initialize the global distribution folder
	configPath := constant.ConfigPath
	if !util.Exists(configPath) {
		if err := os.MkdirAll(configPath, os.ModePerm); err != nil {
			logrus.Errorf("create config folder err: %v", err)
			panic(err)
		}
	}

	configFilePath := constant.ConfigFilePath
	if !util.Exists(configFilePath) {
		file, err := os.Create(configFilePath)
		if err != nil {
			logrus.Errorf("create config.ini err: %v", err)
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
			logrus.Errorf("config.ini file write err: %v", err)
			panic(err)
		}
	}

	sqlitePath := constant.SqlitePath
	if !util.Exists(sqlitePath) {
		if err := os.MkdirAll(sqlitePath, os.ModePerm); err != nil {
			logrus.Errorf("create sqlite folder err: %v", err)
			panic(err)
		}
	}
	sqliteFilePath := constant.SqliteFilePath
	if !util.Exists(sqliteFilePath) {
		file, err := os.Create(sqliteFilePath)
		if err != nil {
			logrus.Errorf("create trojan_panel_core.db err: %v", err)
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

// InitConfig initialize the global configuration file
func InitConfig() {
	if err := ini.MapTo(Config, constant.ConfigFilePath); err != nil {
		logrus.Errorf("configuration file failed to load err: %v", err)
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
	FileName   string `ini:"filename"`    // log file location
	MaxSize    int    `ini:"max_size"`    // the maximum capacity of a single file, in MB
	MaxBackups int    `ini:"max_backups"` // the maximum number of expired files to keep
	MaxAge     int    `ini:"max_age"`     // the maximum time interval for keeping expired files, in days
	Compress   bool   `ini:"compress"`    // do you need to compress the rolling log, using gzip compression
}

// GrpcConfig gRPC
type GrpcConfig struct {
	Port string `ini:"port"` // gRPC port
}

type ServerConfig struct {
	Port int `ini:"port"` // service port
}
