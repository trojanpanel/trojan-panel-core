package service

import (
	"trojan-panel-core/dao"
	"trojan-panel-core/module/bo"
	"trojan-panel-core/module/vo"
)

// SelectAccountByPass 用户认证 hysteria
func SelectAccountByPass(pass string) (*vo.AccountHysteriaVo, error) {
	return dao.SelectAccountByPass(pass)
}

func UpdateAccountFlowByPassOrHash(pass *string, hash *string, download int, upload int) error {
	return dao.UpdateAccountFlowByPassOrHash(pass, hash, download, upload)
}

func SelectAccounts() ([]bo.AccountBo, error) {
	return dao.SelectAccounts()
}
