package dao

import (
	"errors"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"github.com/sirupsen/logrus"
	"trojan-panel-core/module"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

func UpdateAccountById(id int, quota *int, download *int, upload *int) error {
	where := map[string]interface{}{"id": id}
	update := map[string]interface{}{}
	if quota != nil {
		update["quota"] = quota
	}
	if download != nil {
		update["download"] = download
	}
	if upload != nil {
		update["upload"] = upload
	}
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
	return nil
}

// SelectAccount 查询账户 全量
func SelectAccount() ([]module.Account, error) {
	var accounts []module.Account
	selectFields := []string{"id"}
	buildSelect, values, err := builder.BuildSelect(mySQLConfig.AccountTable, nil, selectFields)
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

	if err := scanner.Scan(rows, &accounts); err != nil {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}
	return accounts, nil
}

func BanUsers() ([]module.Account, error) {
	var accounts []module.Account
	data := map[string]interface{}{
		"account_table": mySQLConfig.AccountTable,
		"expire_time":   util.NowMilli,
	}
	buildSelect, values, err := builder.NamedQuery(`select id
from account
where (quota >= 0 and quota <= download + upload)
   or deleted = 1
   or expire_time < {{ expire_time }}`, data)
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

	if err := scanner.Scan(rows, &accounts); err != nil {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}
	return accounts, nil
}
