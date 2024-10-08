package dao

import (
	"errors"
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
	"time"
	"trojan-core/model/constant"
)

var sqlInitStr = "CREATE TABLE IF NOT EXISTS proxy\n(\n    id          INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,\n    proxy_type  INTEGER NOT NULL DEFAULT 0,\n    config      TEXT    NOT NULL DEFAULT '',\n    create_time TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,\n    update_time TIMESTAMP        DEFAULT CURRENT_TIMESTAMP\n)"

var sqliteDB *gorm.DB

func InitSqliteDB() error {
	var err error
	sqliteDB, err = gorm.Open(sqlite.Open(constant.SqliteDBPath), &gorm.Config{
		TranslateError: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				Colorful:                  false,
			},
		),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		logrus.Errorf("sqlite open err: %v", err)
		return errors.New("sqlite open err")
	}

	if err = sqliteInit(sqlInitStr); err != nil {
		return err
	}
	return nil
}

func sqliteInit(sqlStr string) error {
	if sqliteDB != nil {
		sqls := strings.Split(strings.Replace(sqlStr, "\r\n", "\n", -1), ";\n")
		for _, s := range sqls {
			s = strings.TrimSpace(s)
			if s != "" {
				tx := sqliteDB.Exec(s)
				if tx.Error != nil && !strings.HasPrefix(tx.Error.Error(), "SQL logic error: duplicate column name") {
					logrus.Errorf("sqlite exec err: %v", tx.Error)
					return errors.New("sqlite exec err")
				}
			}
		}
	}
	return nil
}

func CloseSqliteDB() error {
	if sqliteDB != nil {
		db, err := sqliteDB.DB()
		if err != nil {
			logrus.Errorf("sqlite err: %v", err)
			return errors.New("sqlite err")
		}
		if err = db.Close(); err != nil {
			logrus.Errorf("sqlite close err: %v", err)
			return errors.New("sqlite close err")
		}
	}
	return nil
}

func Paginate(pageNum *int64, pageSize *int64) func(db *gorm.DB) *gorm.DB {
	var num int64 = 1
	var size int64 = 10
	if pageNum != nil && *pageNum > 0 {
		num = *pageNum
	}
	if pageSize != nil && *pageSize > 0 {
		size = *pageSize
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int((num - 1) * size)).Limit(int(size))
	}
}
