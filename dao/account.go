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
	"trojan-panel-core/util"
)

func UpdateAccountFlowByUsername(username string, download int, upload int) error {
	mySQLConfig := core.Config.MySQLConfig
	where := map[string]interface{}{"username": username}
	update := map[string]interface{}{}
	if download > 0 {
		update["download = download +"] = download
	}
	if upload > 0 {
		update["upload = upload +"] = upload
	}
	if len(update) > 0 {
		buildUpdate, values, err := builder.BuildUpdate(mySQLConfig.AccountTable, where, update)
		if err != nil {
			logrus.Errorln(err.Error())
			return errors.New(constant.SysError)
		}

		_, err = db.Exec(buildUpdate, values...)
		if err != nil {
			logrus.Errorln(err.Error())
			return errors.New(constant.SysError)
		}
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

	data := map[string]interface{}{
		"account_table": mySQLConfig.AccountTable,
	}
	if ban {
		buildSelect, values, err = builder.NamedQuery(
			`select username, pass from {{account_table}} where quota >= 0 and quota <= download + upload`, data)
	} else {
		buildSelect, values, err = builder.NamedQuery(
			`select id, username, pass from {{account_table}} where quota < 0 or quota > download + upload`, data)
	}
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
			passDecode, err := util.AesDecode(*item.Pass)
			if err != nil {
				continue
			}
			password, err := util.AesEncode(fmt.Sprintf("%s%s", *item.Username, passDecode))
			if err != nil {
				continue
			}
			passwords = append(passwords, password)
		}
	}
	return passwords, nil
}

func SelectAccountByUsernameAndPass(username string, pass string) (*vo.AccountHysteriaVo, error) {
	mySQLConfig := core.Config.MySQLConfig
	var account module.Account

	passEncode, err := util.AesEncode(pass)
	if err != nil {
		return nil, err
	}

	buildSelect, values, err := builder.NamedQuery(
		`select id, username from {{account_table}} where quota != 0 and username = {{username}} and pass = {{pass}}`, map[string]interface{}{
			"account_table": mySQLConfig.AccountTable,
			"username":      username,
			"pass":          passEncode,
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
		Id:       *account.Id,
		Username: *account.Username,
	}
	return &AccountHysteriaVo, nil
}
