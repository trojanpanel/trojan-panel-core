package service

import (
	"github.com/sirupsen/logrus"
	"trojan-panel-core/dao"
)

func UpdateAccountById(id int, quota *int, download *int, upload *int) error {
	if err := dao.UpdateAccountById(id, quota, download, upload); err != nil {
		return err
	}
	return nil
}

func CronCalUD() {
	accounts, err := dao.SelectAccountActive()
	if err != nil {
		logrus.Errorf("定时刷新流量 查询全量账户错误 err: %v\n", err)
	}
	for _, item := range accounts {
		download, upload, err := dao.SelectUsersDUByAccountId(*item.Id)
		if err != nil {
			logrus.Errorf("定时刷新流量 查询用户download,upload错误 err: %v\n", err)
			continue
		}
		if err = dao.UpdateAccountById(*item.Id, nil, &download, &upload); err != nil {
			logrus.Errorf("定时刷新流量 更新account download,upload错误 err: %v\n", err)
			continue
		}
	}
}

// CronBanUser 禁用用户
func CronBanUser() {
	users, err := dao.BanUsers()
	if err != nil {
		logrus.Errorf("查询禁用用户错误 err: %v\n", err)
	}
	for _, item := range users {
		if err := dao.UpdateAccountById(*item.Id, new(int), nil, nil); err != nil {
			logrus.Errorf("禁用用户错误 err: %v\n", err)
			continue
		}
		if err := dao.UpdateUser(item.Id, nil, nil, new(int), new(int)); err != nil {
			logrus.Errorf("禁用用户错误 err: %v\n", err)
			continue
		}
	}
}
