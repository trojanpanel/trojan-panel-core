package dao

import (
	"errors"
	"fmt"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"github.com/sirupsen/logrus"
	"trojan-panel-core/core"
	"trojan-panel-core/module"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/vo"
)

func UpdateAccountFlowByPassOrHash(pass *string, hash *string, download int, upload int) error {
	mySQLConfig := core.Config.MySQLConfig

	values := []interface{}{download, upload}
	sql := fmt.Sprintf("update %s set download = download + ?,upload = upload + ? where", mySQLConfig.AccountTable)
	if pass != nil && *pass != "" {
		sql += " pass = ?"
		values = append(values, *pass)
	}
	if hash != nil && *hash != "" {
		sql += " hash = ?"
		values = append(values, *hash)
	}
	_, err := db.Exec(sql, values)
	if err != nil {
		logrus.Errorln(err.Error())
		return errors.New(constant.SysError)
	}
	return nil
}

// SelectAccountPasswords 查询全量账户
func SelectAccountPasswords(ban bool) ([]string, error) {
	mySQLConfig := core.Config.MySQLConfig
	var accounts []module.Account
	var (
		buildSelect string
		values      []interface{}
		err         error
	)

	var where map[string]interface{}
	if ban {
		where = map[string]interface{}{"quota >=": 0, "quota <=": "download + upload"}
	} else {
		where = map[string]interface{}{"_or": []map[string]interface{}{{"quota <": 0}, {"quota >": "download + upload"}}}
	}
	buildSelect, values, err = builder.BuildSelect(mySQLConfig.AccountTable, where, []string{"id", "username", "pass"})
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}
	rows, err := db.Query(buildSelect, values...)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}
	defer rows.Close()

	if err = scanner.Scan(rows, &accounts); err != nil && err != scanner.ErrEmptyResult {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}
	passwords := make([]string, 0)
	if len(accounts) > 0 {
		for _, item := range accounts {
			passwords = append(passwords, *item.Pass)
		}
	}
	return passwords, nil
}

func SelectAccountByPass(pass string) (*vo.AccountHysteriaVo, error) {
	mySQLConfig := core.Config.MySQLConfig
	var account module.Account

	selectFields := []string{"id", "username"}
	where := map[string]interface{}{
		"quota <>": 0,
		"pass":     pass,
	}
	buildSelect, values, err := builder.BuildSelect(mySQLConfig.AccountTable, where, selectFields)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}
	rows, err := db.Query(buildSelect, values...)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}
	defer rows.Close()

	err = scanner.Scan(rows, &account)
	if err == scanner.ErrEmptyResult {
		return nil, errors.New(constant.UsernameOrPassError)
	} else if err != nil {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}

	AccountHysteriaVo := vo.AccountHysteriaVo{
		Id:       *account.Id,
		Username: *account.Username,
	}
	return &AccountHysteriaVo, nil
}
