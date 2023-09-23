package dao

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
	"strings"
	"trojan-panel-core/model/constant"
)

var sqliteDb *sql.DB

// InitSqlLite initialize the sqlite data file
func InitSqlLite() {
	var err error
	sqliteDb, err = sql.Open("sqlite", fmt.Sprintf("file:%s", constant.SqliteFilePath))
	if err != nil {
		logrus.Errorf("sqlite connection err: %v", err)
		panic(err)
	}

	var count uint
	row := sqliteDb.QueryRow("SELECT count(1) FROM sqlite_master WHERE type='table' AND name = 'node_config'")
	if err = row.Scan(&count); err != nil {
		logrus.Errorf("query sqlite database err: %v", err)
		panic(err)
	}
	if count == 0 {
		if err = SqlInit(sqlInitStr); err != nil {
			logrus.Errorf("sqlite database import err: %v", err)
			panic(err)
		}
	}
}

func CloseSqliteDb() {
	if sqliteDb != nil {
		if err := sqliteDb.Close(); err != nil {
			logrus.Errorf("sqlite close err: %v", err)
		}
	}
}

func SqlInit(sqlStr string) error {
	sqls := strings.Split(strings.Replace(sqlStr, "\r\n", "\n", -1), ";\n")
	for _, s := range sqls {
		s = strings.TrimSpace(s)
		if s != "" {
			if _, err := sqliteDb.Exec(s); err != nil {
				logrus.Errorf("sqlite database sql execution err: %v", err)
				return err
			}
		}
	}
	return nil
}

var sqlInitStr = "PRAGMA foreign_keys=OFF;\nBEGIN TRANSACTION;\nCREATE TABLE IF NOT EXISTS \"node_config\"\n(\n    id             INTEGER not null\n        primary key autoincrement,\n    api_port       INTEGER default 0 not null,\n    node_type_id   INTEGER default 0 not null,\n    protocol       TEXT    default '' not null,\n    xray_flow      TEXT    default '' not null,\n    xray_ss_method TEXT    default '' not null\n);\nDELETE FROM sqlite_sequence;\nINSERT INTO sqlite_sequence VALUES('node_config',0);\nCREATE INDEX node_config_api_port_node_type_id_index\n    on node_config (api_port, node_type_id);\nCOMMIT;"
