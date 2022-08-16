package service

import (
	"trojan-panel-core/dao"
	"trojan-panel-core/module"
	"trojan-panel-core/module/vo"
)

// CountUserByApiPort 查询端口是否已用
func CountUserByApiPort(apiPort uint) (int, error) {
	total, err := dao.CountUserByApiPort(apiPort)
	if err != nil {
		return 0, err
	}
	return total, nil
}

// UpdateUser 更新用户
func UpdateUser(accountId *uint, apiPort *uint, password *string, download *int, upload *int) error {
	if err := dao.UpdateUser(accountId, apiPort, password, download, upload); err != nil {
		return err
	}
	return nil
}

// InsertUsers 插入用户
func InsertUsers(users []module.Users) error {
	if err := dao.InsertUsers(users); err != nil {
		return err
	}
	return nil
}

// SelectUsersToApi 查询需要同步至应用的用户
func SelectUsersToApi(isAdd bool) ([]vo.UserApiVo, error) {
	apiUserVos, err := dao.SelectUsersToApi(isAdd)
	if err != nil {
		return nil, err
	}
	return apiUserVos, nil
}
