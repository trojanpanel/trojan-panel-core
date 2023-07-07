package app

import (
	"errors"
	trojangoservice "github.com/p4gefau1t/trojan-go/api/service"
	"github.com/sirupsen/logrus"
	"regexp"
	"sync"
	"trojan-panel-core/app/hysteria"
	"trojan-panel-core/app/naiveproxy"
	"trojan-panel-core/app/trojango"
	"trojan-panel-core/app/xray"
	"trojan-panel-core/core/process"
	"trojan-panel-core/dao"
	"trojan-panel-core/dao/redis"
	"trojan-panel-core/module"
	"trojan-panel-core/module/bo"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/module/vo"
	"trojan-panel-core/service"
	"trojan-panel-core/util"
)

var userLinkRegex = regexp.MustCompile("user>>>([^>]+)>>>traffic>>>(downlink|uplink)")

func InitApp() {
	InitBinFile()
	if err := xray.InitXrayApp(); err != nil {
		logrus.Errorf("Xray app init err: %s", err.Error())
	}
	if err := trojango.InitTrojanGoApp(); err != nil {
		logrus.Errorf("TrojanGo app init err: %s", err.Error())
	}
	if err := hysteria.InitHysteriaApp(); err != nil {
		logrus.Errorf("Hysteria app init err: %s", err.Error())
	}
	if err := naiveproxy.InitNaiveProxyApp(); err != nil {
		logrus.Errorf("NaiveProxy app init err: %s", err.Error())
	}
}

func InitBinFile() {
	if err := xray.InitXrayBinFile(); err != nil {
		logrus.Errorf("Download Xray file err: %v", err)
		panic(err)
	}
	if err := trojango.InitTrojanGoBinFile(); err != nil {
		logrus.Errorf("Download TrojanGo file err: %v", err)
		panic(err)
	}
	if err := hysteria.InitHysteriaBinFile(); err != nil {
		logrus.Errorf("Download Hysteria file err: %v", err)
		panic(err)
	}
	if err := naiveproxy.InitNaiveProxyBinFile(); err != nil {
		logrus.Errorf("Download NaiveProxy file err: %v", err)
		panic(err)
	}
}

func StartApp(nodeAddDto dto.NodeAddDto) error {
	var mutex sync.Mutex
	defer mutex.Unlock()
	if mutex.TryLock() {
		var protocol string
		var xrayFlow string
		var xraySSMethod string
		switch nodeAddDto.NodeTypeId {
		case constant.Xray:
			if err := xray.StartXray(dto.XrayConfigDto{
				ApiPort:        nodeAddDto.Port + 30000,
				Port:           nodeAddDto.Port,
				Protocol:       nodeAddDto.XrayProtocol,
				Settings:       nodeAddDto.XraySettings,
				StreamSettings: nodeAddDto.XrayStreamSettings,
				Tag:            nodeAddDto.XrayTag,
				Sniffing:       nodeAddDto.XraySniffing,
				Allocate:       nodeAddDto.XrayAllocate,
				Template:       nodeAddDto.XrayTemplate,
			}); err != nil {
				return err
			}
			protocol = nodeAddDto.XrayProtocol
			xrayFlow = nodeAddDto.XrayFlow
			xraySSMethod = nodeAddDto.XraySSMethod
		case constant.TrojanGo:
			if err := trojango.StartTrojanGo(dto.TrojanGoConfigDto{
				ApiPort:         nodeAddDto.Port + 30000,
				Port:            nodeAddDto.Port,
				Domain:          nodeAddDto.Domain,
				Sni:             nodeAddDto.TrojanGoSni,
				MuxEnable:       nodeAddDto.TrojanGoMuxEnable,
				WebsocketEnable: nodeAddDto.TrojanGoWebsocketEnable,
				WebsocketPath:   nodeAddDto.TrojanGoWebsocketPath,
				WebsocketHost:   nodeAddDto.TrojanGoWebsocketHost,
				SSEnable:        nodeAddDto.TrojanGoSSEnable,
				SSMethod:        nodeAddDto.TrojanGoSSMethod,
				SSPassword:      nodeAddDto.TrojanGoSSPassword,
			}); err != nil {
				return err
			}
		case constant.Hysteria:
			if err := hysteria.StartHysteria(dto.HysteriaConfigDto{
				ApiPort:  nodeAddDto.Port + 30000,
				Port:     nodeAddDto.Port,
				Protocol: nodeAddDto.HysteriaProtocol,
				Obfs:     nodeAddDto.HysteriaObfs,
				Domain:   nodeAddDto.Domain,
				UpMbps:   nodeAddDto.HysteriaUpMbps,
				DownMbps: nodeAddDto.HysteriaDownMbps,
			}); err != nil {
				return err
			}
		case constant.NaiveProxy:
			if err := naiveproxy.StartNaiveProxy(dto.NaiveProxyConfigDto{
				ApiPort: nodeAddDto.Port + 30000,
				Port:    nodeAddDto.Port,
				Domain:  nodeAddDto.Domain,
			}); err != nil {
				return err
			}
		default:
			return errors.New(constant.NodeTypeNotExist)
		}

		nodeConfig := module.NodeConfig{
			ApiPort:      nodeAddDto.Port + 30000,
			NodeTypeId:   nodeAddDto.NodeTypeId,
			Protocol:     protocol,
			XrayFlow:     xrayFlow,
			XraySSMethod: xraySSMethod,
		}
		if err := service.InsertNodeConfig(nodeConfig); err != nil {
			return err
		}
	}
	return nil
}

func StopApp(apiPort uint, nodeTypeId uint) error {
	var mutex sync.Mutex
	defer mutex.Unlock()
	if mutex.TryLock() {
		switch nodeTypeId {
		case constant.Xray:
			if err := xray.StopXray(apiPort, true); err != nil {
				return err
			}
		case constant.TrojanGo:
			if err := trojango.StopTrojanGo(apiPort, true); err != nil {
				return err
			}
		case constant.Hysteria:
			if err := hysteria.StopHysteria(apiPort, true); err != nil {
				return err
			}
		case constant.NaiveProxy:
			if err := naiveproxy.StopNaiveProxy(apiPort, true); err != nil {
				return err
			}
		default:
			return errors.New(constant.NodeTypeNotExist)
		}

		if err := service.DeleteNodeConfigByNodeTypeIdAndApiPort(apiPort, nodeTypeId); err != nil {
			return err
		}
	}
	return nil
}

func RestartApp(apiPort uint, nodeTypeId uint) error {
	switch nodeTypeId {
	case constant.Xray:
		if err := xray.RestartXray(apiPort); err != nil {
			return err
		}
	case constant.TrojanGo:
		if err := trojango.RestartTrojanGo(apiPort); err != nil {
			return err
		}
	case constant.Hysteria:
		if err := hysteria.RestartHysteria(apiPort); err != nil {
			return err
		}
	case constant.NaiveProxy:
		if err := naiveproxy.RestartNaiveProxy(apiPort); err != nil {
			return err
		}
	default:
		return errors.New(constant.NodeTypeNotExist)
	}
	return nil
}

// CronHandlerUser 定时任务 处理用户
func CronHandlerUser() {
	// 数据库中有效的账户
	accountBos, err := dao.SelectAccounts()
	if err != nil {
		logrus.Errorf("查询全部有效账户错误 err: %v", err)
		return
	}

	// xray
	go func(accountBos []bo.AccountBo) {
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
					if len(submatch) == 3 {
						pass := submatch[1]
						var banFlag = true
						for _, account := range accountBos {
							if account.Pass == pass {
								banFlag = false
								break
							}
						}
						if banFlag {
							banAccountBos = append(banAccountBos, bo.AccountBo{
								Pass: pass,
							})
						}
					}
				}
				for _, item := range banAccountBos {
					if err = xrayApi.DeleteUser(item.Pass); err != nil {
						logrus.Errorf("Xray DeleteUser err: %v", err)
						continue
					}
				}

				// 添加的账户
				var addAccountBos []bo.AccountBo
				for _, account := range accountBos {
					var addFlag = true
					for _, stat := range stats {
						submatch := userLinkRegex.FindStringSubmatch(stat.Name)
						if len(submatch) == 3 {
							pass := submatch[1]
							if account.Pass == pass {
								addFlag = false
								break
							}
						}
					}
					if addFlag {
						addAccountBos = append(addAccountBos, bo.AccountBo{
							Pass: account.Pass,
						})
					}
				}
				protocol, err := util.GetXrayProtocolByApiPort(apiPort.(uint))
				if err != nil {
					logrus.Errorf("Xray get protocol err apiPort: %s err: %v", apiPort, err)
				} else {
					for _, item := range addAccountBos {
						if err = xrayApi.AddUser(dto.XrayAddUserDto{
							Protocol: protocol,
							Password: item.Pass,
						}); err != nil {
							logrus.Errorf("Xray AddUser err: %v", err)
							continue
						}
					}
				}
			}()
			return true
		})
	}(accountBos)

	// trojan go
	go func(accountBos []bo.AccountBo) {
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
					hash := user.GetUser().GetHash()
					var banFlag = true
					for _, account := range accountBos {
						if account.Hash == hash {
							banFlag = false
							break
						}
					}
					if banFlag {
						banAccountBos = append(banAccountBos, bo.AccountBo{
							Hash: hash,
						})
					}
				}
				for _, item := range banAccountBos {
					// 调用api删除用户
					if err = trojanGoApi.DeleteUser(item.Hash); err != nil {
						logrus.Errorf("TrojanGo DeleteUser err: %v", err)
						continue
					}
				}

				// 添加的账户
				var addAccountBos []bo.AccountBo
				for _, account := range accountBos {
					var addFlag = true
					for _, user := range users {
						hash := user.GetUser().GetHash()
						if account.Hash == hash {
							addFlag = false
							break
						}
					}
					if addFlag {
						addAccountBos = append(addAccountBos, bo.AccountBo{
							Hash: account.Hash,
						})
					}
				}
				for _, item := range addAccountBos {
					// 调用api添加用户
					if err = trojanGoApi.AddUser(dto.TrojanGoAddUserDto{
						Hash: item.Hash,
					}); err != nil {
						logrus.Errorf("TrojanGo AddUser err: %v", err)
						continue
					}
				}
			}()
			return true
		})
	}(accountBos)

	// naiveproxy
	go func(accountBos []bo.AccountBo) {
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
					var banFlag = true
					for _, account := range accountBos {
						if account.Pass == pass {
							banFlag = false
							break
						}
					}
					if banFlag {
						banAccountBos = append(banAccountBos, bo.AccountBo{
							Pass: pass,
						})
					}
				}
				for _, item := range banAccountBos {
					// 调用api删除用户
					if err = naiveProxyApi.DeleteUser(item.Pass); err != nil {
						logrus.Errorf("NaiveProxy DeleteUser err: %v", err)
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
							break
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
						logrus.Errorf("NaiveProxy AddUser err: %v", err)
						continue
					}
				}
			}()
			return true
		})
	}(accountBos)
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
					mutex, err := redis.RsLock(constant.LockXrayUpdate)
					if err != nil {
						return
					}
					for _, account := range accountUpdateBos {
						if err = dao.UpdateAccountFlowByPassOrHash(&account.Pass, nil, account.Download, account.Upload); err != nil {
							logrus.Errorf("Xray UpdateAccountFlow err apiPort: %d err: %v", apiPort, err)
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
				go func(users []*trojangoservice.UserStatus) {
					accountUpdateBos := make([]bo.AccountUpdateBo, 0)
					for _, user := range users {
						hash := user.GetUser().GetHash()
						downloadTraffic := int(user.GetTrafficTotal().GetDownloadTraffic())
						uploadTraffic := int(user.GetTrafficTotal().GetUploadTraffic())
						if err = trojanGoApi.ReSetUserTrafficByHash(hash); err != nil {
							logrus.Errorf("TrojanGo ReSetUserTraffic err apiPort: %d err: %v", apiPort, err)
							continue
						}
						accountUpdateBo := bo.AccountUpdateBo{}
						accountUpdateBo.Hash = hash
						accountUpdateBo.Download = downloadTraffic
						accountUpdateBo.Upload = uploadTraffic
						accountUpdateBos = append(accountUpdateBos, accountUpdateBo)
					}

					mutex, err := redis.RsLock(constant.LockTrojanGoUpdate)
					if err != nil {
						return
					}
					for _, account := range accountUpdateBos {
						if err = dao.UpdateAccountFlowByPassOrHash(nil, &account.Hash, account.Download,
							account.Upload); err != nil {
							logrus.Errorf("TrojanGo UpdateAccountFlow err apiPort: %d err: %v", apiPort, err)
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

func RemoveAccount(password string) error {
	// xray
	go func() {
		xrayInstance := process.NewXrayProcess()
		xrayCmdMap := xrayInstance.GetCmdMap()
		xrayCmdMap.Range(func(apiPort, cmd any) bool {
			go func(password string) {
				xrayApi := xray.NewXrayApi(apiPort.(uint))
				if err := xrayApi.DeleteUser(password); err != nil {
					logrus.Errorf("Xray DeleteUser err: %v", err)
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
					logrus.Errorf("TrojanGo DeleteUser err: %v", err)
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
					logrus.Errorf("NavieProxy DeleteUser err: %v", err)
				}
			}(password)
			return true
		})
	}()
	return nil
}
