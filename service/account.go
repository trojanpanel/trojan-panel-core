package service

import "trojan-panel-core/dao"

func UpdateAccountById(id int, quota int) error {
	if err := dao.UpdateAccountById(id, quota); err != nil {
		return err
	}
	return nil
}
