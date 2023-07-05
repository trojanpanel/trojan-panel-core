package service

import (
	"fmt"
	"github.com/p4gefau1t/trojan-go/api/service"
	"github.com/sirupsen/logrus"
	"regexp"
	"trojan-panel-core/app/naiveproxy"
	"trojan-panel-core/app/trojango"
	"trojan-panel-core/app/xray"
	"trojan-panel-core/core/process"
	"trojan-panel-core/dao"
	"trojan-panel-core/dao/redis"
	"trojan-panel-core/module/bo"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/module/vo"
	"trojan-panel-core/util"
)

var userLinkRegex = regexp.MustCompile("user>>>([^>]+)>>>traffic>>>(downlink|uplink)")

// CronHandlerUser 定时任务 处理用户
func CronHandlerUser() {
	// 数据库中有效的账户
	accountBos, err := dao.SelectAccounts()
	if err != nil {
		logrus.Errorf("查询全部有效账户错误 err: %v", err)
		return
	}

	// xray
	go func() {
		xrayInstance := process.NewXrayProcess()
		xrayCmdMap := xrayInstance.GetCmdMap()
		xrayCmdMap.Range(func(apiPort, cmd any) bool {
			go func() {
				xrayApi := xray.NewXrayApi(apiPort.(uint))
				// 内存中的账户
				stats, err := xrayApi.QueryStats("", false)
				if err != nil {
					return
				}

				// 删除的账户
				var banAccountBos []bo.AccountBo
				for _, stat := range stats {
					submatch := userLinkRegex.FindStringSubmatch(stat.Name)
					if len(submatch) == 2 {
						pass := submatch[1]
						var addFlag = true
						for _, account := range accountBos {
							if account.Pass == pass {
								addFlag = false
								continue
							}
						}
						if addFlag {
							banAccountBos = append(banAccountBos, bo.AccountBo{
								Pass: pass,
							})
						}
					}
				}
				for _, item := range banAccountBos {
					if err = xrayApi.DeleteUser(item.Pass); err != nil {
						logrus.Errorf("Xray调用api删除用户错误 err: %v", err)
						continue
					}
				}

				// 添加的账户
				var addAccountBos []bo.AccountBo
				for _, account := range accountBos {
					var addFlag = true
					for _, stat := range stats {
						submatch := userLinkRegex.FindStringSubmatch(stat.Name)
						if len(submatch) == 2 {
							pass := submatch[1]
							if account.Pass == pass {
								addFlag = false
								continue
							}
						}
						if addFlag {
							addAccountBos = append(addAccountBos, bo.AccountBo{
								Pass: account.Pass,
							})
						}
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
			}()
			return true
		})
	}()

	// trojan go
	go func() {
		trojanGoInstance := process.NewTrojanGoInstance()
		trojanGoCmdMap := trojanGoInstance.GetCmdMap()
		trojanGoCmdMap.Range(func(apiPort, cmd any) bool {
			go func() {
				trojanGoApi := trojango.NewTrojanGoApi(apiPort.(uint))
				users, err := trojanGoApi.ListUsers()
				if err != nil {
					return
				}

				// 删除的账户
				var banAccountBos []bo.AccountBo
				for _, user := range users {
					pass := user.GetUser().GetPassword()
					var addFlag = true
					for _, account := range accountBos {
						if account.Pass == pass {
							addFlag = false
							continue
						}
					}
					if addFlag {
						banAccountBos = append(banAccountBos, bo.AccountBo{
							Pass: pass,
						})
					}
				}
				for _, item := range banAccountBos {
					// 调用api删除用户
					if err = trojanGoApi.DeleteUser(item.Pass); err != nil {
						logrus.Errorf("TrojanGo调用api删除用户错误 err: %v", err)
						continue
					}
				}

				// 添加的账户
				var addAccountBos []bo.AccountBo
				for _, account := range accountBos {
					var addFlag = true
					for _, user := range users {
						pass := user.GetUser().GetPassword()
						if account.Pass == pass {
							addFlag = false
							continue
						}
					}
					if addFlag {
						addAccountBos = append(addAccountBos, bo.AccountBo{
							Pass: account.Pass,
						})
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
			}()
			return true
		})
	}()

	// naiveproxy
	go func() {
		naiveProxyInstance := process.NewNaiveProxyInstance()
		naiveProxyCmdMap := naiveProxyInstance.GetCmdMap()
		naiveProxyCmdMap.Range(func(apiPort, cmd any) bool {
			go func() {
				naiveProxyApi := naiveproxy.NewNaiveProxyApi(apiPort.(uint))
				users, err := naiveProxyApi.ListUsers()
				if err != nil {
					return
				}

				// 删除的账户
				var banAccountBos []bo.AccountBo
				for _, user := range *users {
					pass := user.AuthPassDeprecated
					var addFlag = true
					for _, account := range accountBos {
						if account.Pass == pass {
							addFlag = false
							continue
						}
					}
					if addFlag {
						banAccountBos = append(banAccountBos, bo.AccountBo{
							Pass: pass,
						})
					}
				}
				for _, item := range banAccountBos {
					// 调用api删除用户
					if err = naiveProxyApi.DeleteUser(item.Pass); err != nil {
						logrus.Errorf("NaiveProxy调用api删除用户错误 err: %v", err)
						continue
					}
				}

				// 添加的账户
				var addAccountBos []bo.AccountBo
				for _, account := range accountBos {
					var addFlag = true
					for _, user := range *users {
						pass := user.AuthPassDeprecated
						if account.Pass == pass {
							addFlag = false
							continue
						}
					}
					if addFlag {
						addAccountBos = append(addAccountBos, bo.AccountBo{
							Username: account.Username,
							Pass:     account.Pass,
						})
					}
				}
				for _, item := range addAccountBos {
					// 调用api添加用户
					if err = naiveProxyApi.AddUser(dto.NaiveProxyAddUserDto{
						Username: item.Username,
						Pass:     item.Pass,
					}); err != nil {
						logrus.Errorf("NaiveProxy调用api添加用户错误 err: %v", err)
						continue
					}
				}
			}()
			return true
		})
	}()
}

// CronHandlerDownloadAndUpload 定时任务 更新数据库中用户的下载和上传流量 Hysteria暂不支持流量统计
func CronHandlerDownloadAndUpload() {
	// xray
	go func() {
		xrayInstance := process.NewXrayProcess()
		xrayCmdMap := xrayInstance.GetCmdMap()
		xrayCmdMap.Range(func(apiPort, cmd any) bool {
			xrayApi := xray.NewXrayApi(apiPort.(uint))
			stats, err := xrayApi.QueryStats("", true)
			if err == nil {
				go func(stats []vo.XrayStatsVo) {
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
					mutex, err := redis.RsLock(fmt.Sprintf(constant.LockXrayUpdate))
					if err != nil {
						return
					}
					for _, account := range accountUpdateBos {
						if err = dao.UpdateAccountFlowByPassOrHash(&account.Pass, nil, account.Download, account.Upload); err != nil {
							logrus.Errorf("Xray同步至数据库 apiPort: %d 更新用户失败 err: %v", apiPort, err)
							continue
						}
					}
					redis.RsUnLock(mutex)
				}(stats)
			}
			return true
		})
	}()

	// trojango
	go func() {
		trojanGoInstance := process.NewTrojanGoInstance()
		trojanGoCmdMap := trojanGoInstance.GetCmdMap()
		trojanGoCmdMap.Range(func(apiPort, cmd any) bool {
			trojanGoApi := trojango.NewTrojanGoApi(apiPort.(uint))
			users, err := trojanGoApi.ListUsers()
			if err == nil {
				go func(users []*service.UserStatus) {
					accountUpdateBos := make([]bo.AccountUpdateBo, 0)
					for _, user := range users {
						hash := user.GetUser().GetHash()
						downloadTraffic := int(user.GetTrafficTotal().GetDownloadTraffic())
						uploadTraffic := int(user.GetTrafficTotal().GetUploadTraffic())
						if err = trojanGoApi.ReSetUserTrafficByHash(hash); err != nil {
							logrus.Errorf("Trojan Go同步至数据库 apiPort: %d 重设Trojan Go用户流量失败 err: %v", apiPort, err)
							continue
						}
						accountUpdateBo := bo.AccountUpdateBo{}
						accountUpdateBo.Hash = hash
						accountUpdateBo.Download = downloadTraffic
						accountUpdateBo.Upload = uploadTraffic
						accountUpdateBos = append(accountUpdateBos, accountUpdateBo)
					}

					mutex, err := redis.RsLock(fmt.Sprintf(constant.LockTrojanGoUpdate))
					if err != nil {
						return
					}
					for _, account := range accountUpdateBos {
						if err = dao.UpdateAccountFlowByPassOrHash(nil, &account.Hash, account.Download,
							account.Upload); err != nil {
							logrus.Errorf("Trojan Go同步至数据库 apiPort: %d 更新用户失败 err: %v", apiPort, err)
							continue
						}
					}
					redis.RsUnLock(mutex)
				}(users)
			}
			return true
		})
	}()
}

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

func RemoveAccount(password string) error {
	// xray
	go func() {
		xrayInstance := process.NewXrayProcess()
		xrayCmdMap := xrayInstance.GetCmdMap()
		xrayCmdMap.Range(func(apiPort, cmd any) bool {
			go func(password string) {
				xrayApi := xray.NewXrayApi(apiPort.(uint))
				if err := xrayApi.DeleteUser(password); err != nil {
					logrus.Errorf("Xray调用api删除用户错误 err: %v", err)
				}
			}(password)
			return true
		})
	}()

	// trojan go
	go func() {
		trojanGoInstance := process.NewTrojanGoInstance()
		trojanGoCmdMap := trojanGoInstance.GetCmdMap()
		trojanGoCmdMap.Range(func(apiPort, cmd any) bool {
			go func(password string) {
				trojanGoApi := trojango.NewTrojanGoApi(apiPort.(uint))
				// 调用api删除用户
				if err := trojanGoApi.DeleteUser(password); err != nil {
					logrus.Errorf("TrojanGo调用api删除用户错误 err: %v", err)
				}
			}(password)
			return true
		})
	}()

	// naiveproxy
	go func() {
		naiveProxyInstance := process.NewNaiveProxyInstance()
		naiveProxyCmdMap := naiveProxyInstance.GetCmdMap()
		naiveProxyCmdMap.Range(func(apiPort, cmd any) bool {
			go func(password string) {
				naiveProxyApi := naiveproxy.NewNaiveProxyApi(apiPort.(uint))
				// 调用api删除用户
				if err := naiveProxyApi.DeleteUser(password); err != nil {
					logrus.Errorf("NavieProxy调用api删除用户错误 err: %v", err)
				}
			}(password)
			return true
		})
	}()
	return nil
}
