package service

import (
	"github.com/sirupsen/logrus"
	"regexp"
	"strings"
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

// CronHandlerDownloadAndUpload 定时任务 更新数据库中用户的下载和上传流量
func CronHandlerDownloadAndUpload() {
	trojanGoInstance := process.NewTrojanGoInstance()
	xrayInstance := process.NewXrayProcess()
	trojanGoCmdMaps := trojanGoInstance.GetCmdMap()
	xrayCmdMaps := xrayInstance.GetCmdMap()

	trojanGoCmdMaps.Range(func(apiPort, cmd any) bool {
		trojanGoApi := trojango.NewTrojanGoApi(apiPort.(uint))
		users, err := trojanGoApi.ListUsers()
		if err == nil {
			for _, user := range users {
				downloadTraffic := int(user.GetTrafficTotal().GetDownloadTraffic())
				uploadTraffic := int(user.GetTrafficTotal().GetDownloadTraffic())
				password := user.GetUser().GetPassword()
				passwordSplit := strings.Split(password, "&")
				if len(passwordSplit) != 2 || len(passwordSplit[0]) == 0 {
					continue
				}
				if err := dao.UpdateAccountFlowByUsername(passwordSplit[0], downloadTraffic,
					uploadTraffic); err != nil {
					logrus.Errorf("Trojan Go同步至数据库 apiPort: %d 更新用户失败 err: %v", apiPort, err)
					continue
				}
			}
		}
		return true
	})

	xrayCmdMaps.Range(func(apiPort, cmd any) bool {
		xrayApi := xray.NewXrayApi(apiPort.(uint))
		stats, err := xrayApi.QueryStats("", true)
		if err != nil {
			logrus.Errorf("数据库同步至Xray apiPort: %d 查询用户失败 err: %v\n", apiPort, err)
		} else {
			accountUpdateBos := make([]bo.AccountUpdateBo, 0)
			for _, stat := range stats {
				submatch := userLinkRegex.FindStringSubmatch(stat.Name)
				accountUpdateBo := bo.AccountUpdateBo{}
				if len(submatch) > 0 {
					email := submatch[0]
					isDown := submatch[1] == "downlink"
					emailSplit := strings.Split(email, "&")
					if len(emailSplit) != 2 || len(emailSplit[0]) == 0 {
						continue
					}
					accountUpdateBo.Username = emailSplit[0]
					if isDown {
						accountUpdateBo.Download = stat.Value
					} else {
						accountUpdateBo.Upload = stat.Value
					}
					accountUpdateBos = append(accountUpdateBos, accountUpdateBo)
				}
			}
			for _, account := range accountUpdateBos {
				if err := dao.UpdateAccountFlowByUsername(account.Username, account.Download, account.Upload); err != nil {
					logrus.Errorf("Xray同步至数据库 apiPort: %d 更新用户失败 err: %v", apiPort, err)
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
