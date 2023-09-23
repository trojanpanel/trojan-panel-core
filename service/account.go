package service

import (
	"trojan-panel-core/dao"
	"trojan-panel-core/model/bo"
	"trojan-panel-core/model/vo"
)

// SelectAccountByPass hysteria account authentication
func SelectAccountByPass(pass string) (*vo.AccountHysteriaVo, error) {
	return dao.SelectAccountByPass(pass)
}

func UpdateAccountFlowByPassOrHash(pass *string, hash *string, download int, upload int) error {
	return dao.UpdateAccountFlowByPassOrHash(pass, hash, download, upload)
}

func SelectAccounts() ([]bo.AccountBo, error) {
	return dao.SelectAccounts()
}
