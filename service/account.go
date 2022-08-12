package service

import "trojan-panel-core/dao"

func UpdateAccountById(id int, download int, upload int) error {
	if err := dao.UpdateAccountById(id, download, upload); err != nil {
		return err
	}
	return nil
}
