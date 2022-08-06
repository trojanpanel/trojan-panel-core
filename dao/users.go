package dao

import (
	"errors"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"github.com/sirupsen/logrus"
	"trojan-panel-core/core"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

var mySQLConfig = core.Config.MySQLConfig

// SelectUsersPassword 查询所有用户 用于api全量更新用户
func SelectUsersPassword() ([]string, error) {
	var passwords []string

	buildSelect, values, err := builder.BuildSelect(mySQLConfig.Table, nil, []string{"`password`"})
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

	if err := scanner.Scan(rows, &passwords); err != nil {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}
	return passwords, nil
}

// UpdateUserDownloadAndUpload 更新用户的download和upload字段
func UpdateUserDownloadAndUpload(password string, download int, upload int) error {
	passwordEncode, err := util.AesEncode(password)
	if err != nil {
		return err
	}
	where := map[string]interface{}{"password": passwordEncode}
	update := map[string]interface{}{
		"download": download,
		"upload":   upload,
	}

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
	return nil
}

// DeleteUserByPassword 根据密码删除用户，用于封禁用户的情况
func DeleteUserByPassword(password string) error {
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

// DeleteUsersByQuota 删除总流量大于配额的情况
func DeleteUsersByQuota() error {
	buildDelete, values, err := builder.BuildDelete(mySQLConfig.Table, map[string]interface{}{
		"quota <=": "download + upload",
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
