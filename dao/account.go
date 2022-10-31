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
	if download == 0 && upload == 0 {
		return nil
	}

	mySQLConfig := core.Config.MySQLConfig

	values := []interface{}{download, upload}
	downloadUpdateSql := ""
	if download != 0 {
		downloadUpdateSql = "download = download + ?"
	}
	uploadUpdateSql := ""
	if upload != 0 {
		if downloadUpdateSql == "" {
			uploadUpdateSql = "upload = upload + ?"
		} else {
			uploadUpdateSql = ",upload = upload + ?"
		}
	}

	sql := fmt.Sprintf("update %s set %s where", mySQLConfig.AccountTable, downloadUpdateSql+uploadUpdateSql)

	if pass != nil && *pass != "" {
		sql += " pass = ?"
		values = append(values, *pass)
	}
	if hash != nil && *hash != "" {
		sql += " hash = ?"
		values = append(values, *hash)
	}
	_, err := db.Exec(sql, values...)
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
		values []interface{}
		err    error
	)

	sql := fmt.Sprintf("select id,pass from %s where", mySQLConfig.AccountTable)
	if ban {
		sql += " quota = 0 or (quota > 0 and quota <= download + upload)"
	} else {
		sql += " quota < 0 or (quota > download + upload)"
	}
	rows, err := db.Query(sql, values...)
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

	selectFields := []string{"id"}
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
		Id: *account.Id,
	}
	return &AccountHysteriaVo, nil
}
