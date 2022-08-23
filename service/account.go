package service

import (
	"github.com/sirupsen/logrus"
	"trojan-panel-core/app/trojango"
	"trojan-panel-core/app/xray"
	"trojan-panel-core/core/process"
	"trojan-panel-core/dao"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/module/vo"
	"trojan-panel-core/util"
)

// CronHandlerUser 定时任务 处理用户
func CronHandlerUser() {
	// 禁用用户
	banPasswords, err := dao.SelectAccountPasswords(true)
	if err != nil {
		logrus.Errorf("查询禁用用户错误 err: %v\n", err)
		return
	}

	// 添加用户
	addPasswords, err := dao.SelectAccountPasswords(false)
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
		for _, item := range banPasswords {
			// 调用api删除用户
			if err = trojanGoApi.DeleteUser(item); err != nil {
				logrus.Errorf("调用api删除用户错误 err: %v\n", err)
				continue
			}
		}

		for _, item := range addPasswords {
			// 调用api添加用户
			if err = trojanGoApi.AddUser(dto.TrojanGoAddUserDto{
				Password: item,
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
		for _, item := range banPasswords {
			if err = xrayApi.DeleteUser(item); err != nil {
				logrus.Errorf("调用api删除用户错误 err: %v\n", err)
				continue
			}
		}
		protocol, err := util.GetXrayProtocolByApiPort(apiPort.(uint))
		if err == nil {
			for _, item := range addPasswords {
				if err := xrayApi.AddUser(dto.XrayAddUserDto{
					Protocol: protocol,
					Password: item,
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
func SelectAccountByUsernameAndPass(username string, pass string) (*vo.AccountHysteriaVo, error) {
	return dao.SelectAccountByUsernameAndPass(username, pass)
}

func UpdateAccountFlowByUsername(username string, download int, upload int) error {
	return dao.UpdateAccountFlowByUsername(username, download, upload)
}

func SelectAccountPasswords(ban bool) ([]string, error) {
	return dao.SelectAccountPasswords(ban)
}
