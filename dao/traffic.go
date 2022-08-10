package dao

import (
	"errors"
	"github.com/didi/gendry/builder"
	"github.com/sirupsen/logrus"
	"trojan-panel-core/module"
	"trojan-panel-core/module/constant"
)

func CountTrafficByApiPort(apiPort int) (int, error) {
	var total int
	selectFields := []string{"count(1)"}
	where := map[string]interface{}{"api_port": apiPort}
	buildSelect, values, err := builder.BuildSelect(mySQLConfig.TrafficTable, where, selectFields)
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

func UpdateTrafficByApiPort(apiPort int, download int, upload int) error {
	where := map[string]interface{}{"api_port": apiPort}
	update := map[string]interface{}{"download": download, "upload": upload}
	buildUpdate, values, err := builder.BuildUpdate(mySQLConfig.TrafficTable, where, update)
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

func SaveTraffic(traffics []module.Traffic) error {
	var data []map[string]interface{}
	for _, item := range traffics {
		traffic := map[string]interface{}{
			"api_port": item.ApiPort,
			"download": item.Download,
			"upload":   item.Upload,
		}
		data = append(data, traffic)
	}
	buildInsert, values, err := builder.BuildInsert(mySQLConfig.TrafficTable, data)
	if err != nil {
		logrus.Errorln(err.Error())
		return errors.New(constant.SysError)
	}
	if _, err := db.Exec(buildInsert, values...); err != nil {
		logrus.Errorln(err.Error())
		return errors.New(constant.SysError)
	}
	return nil
}
