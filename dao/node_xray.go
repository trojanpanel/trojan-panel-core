package dao

import (
	"errors"
	"github.com/didi/gendry/builder"
	"github.com/sirupsen/logrus"
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
