package service

import (
	"github.com/sirupsen/logrus"
	"regexp"
	"trojan-panel-core/app/trojango"
	"trojan-panel-core/app/xray"
	"trojan-panel-core/core/process"
	"trojan-panel-core/dao"
	"trojan-panel-core/module/bo"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/module/vo"
	"trojan-panel-core/util"
)

var userLinkRegex = regexp.MustCompile("user>>>([^>]+)>>>traffic>>>(downlink|uplink)")

// CronHandlerUser 定时任务 处理用户
func CronHandlerUser() {
	// 禁用用户
	banAccountBos, err := dao.SelectAccounts(true)
	if err != nil {
		logrus.Errorf("查询禁用用户错误 err: %v", err)
		return
	}

	// 添加用户
	addAccountBos, err := dao.SelectAccounts(false)
	if err != nil {
		logrus.Errorf("查询添加用户错误 err: %v", err)
		return
	}

	trojanGoInstance := process.NewTrojanGoInstance()
	trojanGoCmdMaps := trojanGoInstance.GetCmdMap()
	xrayInstance := process.NewXrayProcess()
	xrayCmdMaps := xrayInstance.GetCmdMap()

	// xray
	xrayCmdMaps.Range(func(apiPort, cmd any) bool {
		xrayApi := xray.NewXrayApi(apiPort.(uint))
		for _, item := range banAccountBos {
			if err = xrayApi.DeleteUser(item.Pass); err != nil {
				logrus.Errorf("Xray调用api删除用户错误 err: %v", err)
				continue
			}
		}
		protocol, err := util.GetXrayProtocolByApiPort(apiPort.(uint))
		if err != nil {
			logrus.Errorf("Xray查询协议错误 apiPort: %s err: %v", apiPort, err)
		} else {
			for _, item := range addAccountBos {
				if err = xrayApi.AddUser(dto.XrayAddUserDto{
					Protocol: protocol,
					Password: item.Pass,
				}); err != nil {
					logrus.Errorf("Xray调用api添加用户错误 err: %v", err)
					continue
				}
			}
		}
		return true
	})

	// trojan go
	trojanGoCmdMaps.Range(func(apiPort, cmd any) bool {
		trojanGoApi := trojango.NewTrojanGoApi(apiPort.(uint))
		for _, item := range banAccountBos {
			// 调用api删除用户
			if err = trojanGoApi.DeleteUser(item.Pass); err != nil {
				logrus.Errorf("TrojanGo调用api删除用户错误 err: %v", err)
				continue
			}
		}

		for _, item := range addAccountBos {
			// 调用api添加用户
			if err = trojanGoApi.AddUser(dto.TrojanGoAddUserDto{
				Password: item.Pass,
			}); err != nil {
				logrus.Errorf("TrojanGo调用api添加用户错误 err: %v", err)
				continue
			}
		}
		return true
	})
}

// CronHandlerDownloadAndUpload 定时任务 更新数据库中用户的下载和上传流量 Hysteria暂不支持流量统计
func CronHandlerDownloadAndUpload() {
	xrayInstance := process.NewXrayProcess()
	xrayCmdMaps := xrayInstance.GetCmdMap()
	trojanGoInstance := process.NewTrojanGoInstance()
	trojanGoCmdMaps := trojanGoInstance.GetCmdMap()

	xrayCmdMaps.Range(func(apiPort, cmd any) bool {
		xrayApi := xray.NewXrayApi(apiPort.(uint))
		stats, err := xrayApi.QueryStats("", true)
		if err != nil {
			logrus.Errorf("数据库同步至Xray apiPort: %d 查询用户失败 err: %v", apiPort, err)
		} else {
			accountUpdateBos := make([]bo.AccountUpdateBo, 0)
			for _, stat := range stats {
				submatch := userLinkRegex.FindStringSubmatch(stat.Name)
				if len(submatch) == 3 {
					isDown := submatch[2] == "downlink"
					var setFlag = false
					if isDown {
						for index := range accountUpdateBos {
							if accountUpdateBos[index].Pass == submatch[1] {
								accountUpdateBos[index].Download = stat.Value
								setFlag = true
								break
							}
						}
						if !setFlag {
							accountUpdateBo := bo.AccountUpdateBo{}
							accountUpdateBo.Pass = submatch[1]
							accountUpdateBo.Download = stat.Value
							accountUpdateBos = append(accountUpdateBos, accountUpdateBo)
						}
					} else {
						for index := range accountUpdateBos {
							if accountUpdateBos[index].Pass == submatch[1] {
								accountUpdateBos[index].Upload = stat.Value
								setFlag = true
								break
							}
						}
						if !setFlag {
							accountUpdateBo := bo.AccountUpdateBo{}
							accountUpdateBo.Pass = submatch[1]
							accountUpdateBo.Upload = stat.Value
							accountUpdateBos = append(accountUpdateBos, accountUpdateBo)
						}
					}
				}
			}
			for _, account := range accountUpdateBos {
				if err = dao.UpdateAccountFlowByPassOrHash(&account.Pass, nil, account.Download, account.Upload); err != nil {
					logrus.Errorf("Xray同步至数据库 apiPort: %d 更新用户失败 err: %v", apiPort, err)
					continue
				}
			}
		}
		return true
	})

	trojanGoCmdMaps.Range(func(apiPort, cmd any) bool {
		trojanGoApi := trojango.NewTrojanGoApi(apiPort.(uint))
		users, err := trojanGoApi.ListUsers()
		if err == nil {
			for _, user := range users {
				downloadTraffic := int(user.GetTrafficTotal().GetDownloadTraffic())
				uploadTraffic := int(user.GetTrafficTotal().GetUploadTraffic())
				hash := user.GetUser().GetHash()
				if err = trojanGoApi.ReSetUserTrafficByHash(hash); err != nil {
					logrus.Errorf("Trojan Go同步至数据库 apiPort: %d 重设Trojan Go用户流量失败 err: %v", apiPort, err)
					continue
				}
				if err = dao.UpdateAccountFlowByPassOrHash(nil, &hash, downloadTraffic,
					uploadTraffic); err != nil {
					logrus.Errorf("Trojan Go同步至数据库 apiPort: %d 更新用户失败 err: %v", apiPort, err)
					continue
				}
			}
		}
		return true
	})
}

// SelectAccountByPass 用户认证 hysteria
func SelectAccountByPass(pass string) (*vo.AccountHysteriaVo, error) {
	return dao.SelectAccountByPass(pass)
}

func UpdateAccountFlowByPassOrHash(pass *string, hash *string, download int, upload int) error {
	return dao.UpdateAccountFlowByPassOrHash(pass, hash, download, upload)
}

func SelectAccounts(ban bool) ([]bo.AccountBo, error) {
	return dao.SelectAccounts(ban)
}

func RemoveAccount(password string) error {
	xrayInstance := process.NewXrayProcess()
	xrayCmdMaps := xrayInstance.GetCmdMap()
	trojanGoInstance := process.NewTrojanGoInstance()
	trojanGoCmdMaps := trojanGoInstance.GetCmdMap()

	// xray
	xrayCmdMaps.Range(func(apiPort, cmd any) bool {
		xrayApi := xray.NewXrayApi(apiPort.(uint))
		if err := xrayApi.DeleteUser(password); err != nil {
			logrus.Errorf("调用api删除用户错误 err: %v", err)
		}
		return true
	})

	// trojan go
	trojanGoCmdMaps.Range(func(apiPort, cmd any) bool {
		trojanGoApi := trojango.NewTrojanGoApi(apiPort.(uint))
		// 调用api删除用户
		if err := trojanGoApi.DeleteUser(password); err != nil {
			logrus.Errorf("调用api删除用户错误 err: %v", err)
		}
		return true
	})
	return nil
}
