package service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"trojan-panel-core/app/trojango"
	"trojan-panel-core/app/xray"
	"trojan-panel-core/core/process"
	"trojan-panel-core/dao"
	"trojan-panel-core/module"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/module/vo"
	"trojan-panel-core/util"
)

// CronHandlerUser 定时任务 处理用户
func CronHandlerUser() {
	// 禁用用户
	banAccounts, err := dao.SelectAccounts(true)
	if err != nil {
		logrus.Errorf("查询禁用用户错误 err: %v\n", err)
		return
	}

	// 添加用户
	addAccounts, err := dao.SelectAccounts(false)
	if err != nil {
		logrus.Errorf("查询添加用户错误 err: %v\n", err)
		return
	}

	trojanGoInstance := process.NewTrojanGoInstance()
	xrayInstance := process.NewXrayProcess()
	trojanGoCmdMaps := trojanGoInstance.GetCmdMap()
	xrayCmdMaps := xrayInstance.GetCmdMap()

	// trojan go
	trojanGoCmdMaps.Range(func(apiPort, cmd any) bool {
		trojanGoApi := trojango.NewTrojanGoApi(apiPort.(uint))
		for _, item := range banAccounts {
			// 调用api删除用户
			if err = trojanGoApi.DeleteUser(fmt.Sprintf("%s&%s", *item.Username, *item.Pass)); err != nil {
				logrus.Errorf("调用api删除用户错误 err: %v\n", err)
				continue
			}
		}

		for _, item := range addAccounts {
			// 调用api添加用户
			if err = trojanGoApi.AddUser(dto.TrojanGoAddUserDto{
				Password: fmt.Sprintf("%s&%s", *item.Username, *item.Pass),
			}); err != nil {
				logrus.Errorf("调用api添加用户错误 err: %v\n", err)
				continue
			}
		}
		return true
	})

	// xray
	xrayCmdMaps.Range(func(apiPort, cmd any) bool {
		xrayApi := xray.NewXrayApi(apiPort.(uint))
		for _, item := range banAccounts {
			if err = xrayApi.DeleteUser(fmt.Sprintf("%s&%s", *item.Username, *item.Pass)); err != nil {
				logrus.Errorf("调用api删除用户错误 err: %v\n", err)
				continue
			}
		}
		protocol, err := util.GetXrayProtocolByApiPort(apiPort.(uint))
		if err == nil {
			for _, item := range addAccounts {
				if err := xrayApi.AddUser(dto.XrayAddUserDto{
					Protocol: protocol,
					Password: fmt.Sprintf("%s&%s", *item.Username, *item.Pass),
				}); err != nil {
					logrus.Errorf("调用api添加用户错误 err: %v\n", err)
					continue
				}
			}
		}
		return true
	})
}

// SelectAccountByUsernameAndPass 用户认证 hysteria
func SelectAccountByUsernameAndPass(username string, pass string) (*vo.AccountVo, error) {
	accountVo, err := dao.SelectAccountByUsernameAndPass(username, pass)
	if err != nil {
		return nil, err
	}
	return accountVo, nil
}

func UpdateAccountFlowByUsername(username string, download int, upload int) error {
	if err := dao.UpdateAccountFlowByUsername(username, download, upload); err != nil {
		return err
	}
	return nil
}

func SelectAccounts(ban bool) ([]module.Account, error) {
	return dao.SelectAccounts(ban)
}
