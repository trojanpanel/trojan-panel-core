package dao

import (
	"database/sql"
	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"net/url"
	"time"
	"trojan-panel-core/core"
)

var db *sql.DB

// InitMySQL 初始化数据库
func InitMySQL() {
	mySQLConfig := core.Config.MySQLConfig
	var err error

	db, err = manager.
		New(mySQLConfig.Database, mySQLConfig.User, mySQLConfig.Password, mySQLConfig.Host).
		Set(
			manager.SetCharset("utf8"),
			manager.SetAllowCleartextPasswords(true),
			manager.SetInterpolateParams(true),
			manager.SetTimeout(1*time.Second),
			manager.SetReadTimeout(1*time.Second),
			manager.SetLoc(url.QueryEscape("Asia/Shanghai"))).
		Port(mySQLConfig.Port).Open(true)

	if err != nil {
		logrus.Errorf("数据库连接异常 err: %v\n", err)
		panic(err)
	}
}
