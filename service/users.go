package service

import (
	"trojan-panel-core/dao"
	"trojan-panel-core/module/vo"
)

func SelectUsersPassword(isAdd bool) ([]vo.ApiUserVo, error) {
	apiUserVos, err := dao.SelectUsersPassword(isAdd)
	if err != nil {
		return nil, err
	}
	return apiUserVos, nil
}

func UpdateUser(password string, download *int, upload *int, quota *int) error {
	if err := dao.UpdateUser(password, download, upload, quota); err != nil {
		return err
	}
	return nil
}
