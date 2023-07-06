package dao

import (
	"errors"
	"fmt"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"github.com/sirupsen/logrus"
	"trojan-panel-core/core"
	"trojan-panel-core/module"
	"trojan-panel-core/module/bo"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/vo"
)

func UpdateAccountFlowByPassOrHash(pass *string, hash *string, download int, upload int) error {
	if download == 0 && upload == 0 {
		return nil
	}

	mySQLConfig := core.Config.MySQLConfig

	var values []interface{}
	downloadUpdateSql := ""
	if download != 0 {
		downloadUpdateSql = "download = download + ?"
		values = append(values, download)
	}
	uploadUpdateSql := ""
	if upload != 0 {
		if downloadUpdateSql == "" {
			uploadUpdateSql = "upload = upload + ?"
		} else {
			uploadUpdateSql = ",upload = upload + ?"
		}
		values = append(values, upload)
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

// SelectAccounts 查询全量有效的账户
func SelectAccounts() ([]bo.AccountBo, error) {
	mySQLConfig := core.Config.MySQLConfig
	var accounts []module.Account
	var (
		values []interface{}
		err    error
	)

	sql := fmt.Sprintf("select id,username,pass,hash from %s where quota < 0 or (quota > download + upload)", mySQLConfig.AccountTable)
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
	accountBos := make([]bo.AccountBo, 0)
	if len(accounts) > 0 {
		for _, item := range accounts {
			accountBo := bo.AccountBo{
				Username: *item.Username,
				Pass:     *item.Pass,
				Hash:     *item.Hash,
			}
			accountBos = append(accountBos, accountBo)
		}
	}
	return accountBos, nil
}

func SelectAccountByPass(pass string) (*vo.AccountHysteriaVo, error) {
	mySQLConfig := core.Config.MySQLConfig
	var account module.Account

	buildSelect, values, err := builder.NamedQuery(fmt.Sprintf("select id from %s where (quota < 0 or quota > download + upload) and pass = {{pass}}", mySQLConfig.AccountTable),
		map[string]interface{}{
			"pass": pass,
		})
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
