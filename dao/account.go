package dao

import (
	"errors"
	"github.com/didi/gendry/builder"
	"github.com/sirupsen/logrus"
	"trojan-panel-core/module/constant"
)

func UpdateAccountById(id int, quota int) error {
	where := map[string]interface{}{"id": id}
	update := map[string]interface{}{"quota": quota}
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
