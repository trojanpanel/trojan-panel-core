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
	"trojan-panel-core/util"
)

var mySQLConfig = core.Config.MySQLConfig

func SelectUsersXrayPasswords() ([]string, error) {
	var usersXrays []module.UsersXray

	buildSelect, values, err := builder.NamedQuery(
		fmt.Sprintf("select `password` from %s", mySQLConfig.Table), nil)
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

	if err := scanner.Scan(rows, &usersXrays); err != nil {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}

	var passwords []string
	for _, item := range usersXrays {
		passwordDecode, err := util.AesDecode(*item.Password)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, passwordDecode)
	}
	return passwords, nil
}

func UpdateUsersXray(usersXray *module.UsersXray) error {
	passwordEncode, err := util.AesEncode(*usersXray.Password)
	if err != nil {
		return err
	}
	where := map[string]interface{}{"password": passwordEncode}
	update := map[string]interface{}{}
	if usersXray.Download != nil {
		update["download"] = *usersXray.Download
	}
	if usersXray.Upload != nil {
		update["upload"] = *usersXray.Upload
	}

	if len(update) > 0 {
		buildUpdate, values, err := builder.BuildUpdate(mySQLConfig.Table, where, update)
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

// DeleteUsersXray 根据密码删除用户，用于封禁用户的情况
func DeleteUsersXray(password string) error {
	passwordEncode, err := util.AesEncode(password)
	if err != nil {
		return err
	}
	buildDelete, values, err := builder.BuildDelete(mySQLConfig.Table, map[string]interface{}{"password": passwordEncode})
	if err != nil {
		logrus.Errorln(err.Error())
		return errors.New(constant.SysError)
	}

	if _, err := db.Exec(buildDelete, values...); err != nil {
		logrus.Errorln(err.Error())
		return errors.New(constant.SysError)
	}
	return nil
}

// DeleteUsersXrayByQuota 删除总流量大于配额的情况
func DeleteUsersXrayByQuota() error {
	buildDelete, values, err := builder.BuildDelete(mySQLConfig.Table, map[string]interface{}{
		"quota <":  "download + upload",
		"quota >=": 0})
	if err != nil {
		logrus.Errorln(err.Error())
		return errors.New(constant.SysError)
	}

	if _, err := db.Exec(buildDelete, values...); err != nil {
		logrus.Errorln(err.Error())
		return errors.New(constant.SysError)
	}
	return nil
}
