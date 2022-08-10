package service

import (
	"trojan-panel-core/dao"
	"trojan-panel-core/module"
	"trojan-panel-core/module/vo"
)

func CountUserByApiPort(apiPort int) (int, error) {
	total, err := dao.CountUserByApiPort(apiPort)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func UpdateUser(password string, download *int, upload *int) error {
	if err := dao.UpdateUser(password, download, upload); err != nil {
		return err
	}
	return nil
}

func InsertUsers(users []module.Users) error {
	if err := dao.InsertUsers(users); err != nil {
		return err
	}
	return nil
}

func SelectUsersToApi(isAdd bool) ([]vo.ApiUserVo, error) {
	apiUserVos, err := dao.SelectUsersToApi(isAdd)
	if err != nil {
		return nil, err
	}
	return apiUserVos, nil
}
