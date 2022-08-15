package service

import (
	"github.com/sirupsen/logrus"
	"trojan-panel-core/dao"
)

func UpdateAccountById(id int, download int, upload int) error {
	if err := dao.UpdateAccountById(id, download, upload); err != nil {
		return err
	}
	return nil
}

func CronCalUD() {
	accounts, err := dao.SelectAccount()
	if err != nil {
		logrus.Errorf("定时刷新流量 查询全量账户错误 err: %v\n", err)
	}
	for _, item := range accounts {
		download, upload, err := dao.SelectUsersDUByAccountId(*item.Id)
		if err != nil {
			logrus.Errorf("定时刷新流量 查询用户download,upload错误 err: %v\n", err)
			continue
		}
		if err = dao.UpdateAccountById(*item.Id, download, upload); err != nil {
			logrus.Errorf("定时刷新流量 更新account download,upload错误 err: %v\n", err)
			continue
		}
	}
}
