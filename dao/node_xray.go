package dao

import (
	"errors"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	"github.com/sirupsen/logrus"
	"trojan-panel-core/module"
	"trojan-panel-core/module/constant"
)

func CountNodeType() (int, error) {
	var total int

	selectFields := []string{"count(1)"}
	buildSelect, values, err := builder.BuildSelect(mySQLConfig.NodeXray, nil, selectFields)
	if err != nil {
		logrus.Errorln(err.Error())
		return 0, errors.New(constant.SysError)
	}
	if err = db.QueryRow(buildSelect, values...).Scan(&total); err != nil {
		logrus.Errorln(err.Error())
		return 0, errors.New(constant.SysError)
	}
	return total, nil
}

func SelectNodeXrayByProtocol() ([]module.NodeXray, error) {
	var nodeXrays []module.NodeXray
	selectFields := []string{"id", "protocol", "ss_method", "vless_id", "vmess_id", "vmess_alter_id"}
	buildSelect, values, err := builder.BuildSelect(mySQLConfig.NodeXray, nil, selectFields)
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

	err = scanner.Scan(rows, &nodeXrays)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil, errors.New(constant.SysError)
	}
	return nodeXrays, nil
}
