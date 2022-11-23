package naiveproxy

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
	"trojan-panel-core/core"
	"trojan-panel-core/core/process"
	"trojan-panel-core/dao"
	"trojan-panel-core/module/bo"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/util"
)

// InitNaiveProxyApp 初始化NaiveProxy应用
func InitNaiveProxyApp() error {
	apiPorts, err := util.GetConfigApiPorts(constant.NaiveProxyBinPath)
	if err != nil {
		return err
	}
	naiveProxyInstance := process.NewNaiveProxyInstance()
	for _, apiPort := range apiPorts {
		if err = naiveProxyInstance.StartNaiveProxy(apiPort); err != nil {
			return err
		}
	}
	return nil
}

// StartNaiveProxy 启动NaiveProxy
func StartNaiveProxy(naiveProxyConfigDto dto.NaiveProxyConfigDto) error {
	var err error
	if err = initNaiveProxy(naiveProxyConfigDto); err != nil {
		return err
	}
	if err = process.NewNaiveProxyInstance().StartNaiveProxy(naiveProxyConfigDto.ApiPort); err != nil {
		return err
	}
	return nil
}

// StopNaiveProxy 暂停NaiveProxy
func StopNaiveProxy(apiPort uint, removeFile bool) error {
	if err := process.NewNaiveProxyInstance().Stop(apiPort, removeFile); err != nil {
		logrus.Errorf("naiveproxy stop err: %v", err)
		return err
	}
	return nil
}

// RestartNaiveProxy 重启NaiveProxy
func RestartNaiveProxy(apiPort uint) error {
	if err := StopNaiveProxy(apiPort, false); err != nil {
		return err
	}
	if err := StartNaiveProxy(dto.NaiveProxyConfigDto{ApiPort: apiPort}); err != nil {
		return err
	}
	return nil
}

// 初始化NaiveProxy文件
func initNaiveProxy(naiveProxyConfigDto dto.NaiveProxyConfigDto) error {
	// 初始化配置
	naiveProxyConfigFilePath, err := util.GetConfigFilePath(constant.NaiveProxy, naiveProxyConfigDto.ApiPort)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(naiveProxyConfigFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		logrus.Errorf("创建naiveproxy %s文件异常 err: %v", naiveProxyConfigFilePath, err)
		return err
	}
	defer file.Close()

	certConfig := core.Config.CertConfig
	configContent := `{
    "admin": {
        "disabled": true
    },
    "logging": {
        "sink": {
            "writer": {
                "output": "discard"
            }
        },
        "logs": {
            "default": {
                "writer": {
                    "output": "discard"
                }
            }
        }
    },
    "apps": {
        "http": {
            "servers": {
                "srv0": {
                    "listen": [
                        ":${port}"
                    ],
                    "routes": [
                        {
                            "handle": [
                                {
                                    "handler": "subroute",
                                    "routes": [
                                        {
                                            "match": [
                                                {
                                                    "host": [
                                                        "${ip}"
                                                    ]
                                                }
                                            ],
                                            "handle": [
                                                {
                                                    "handler": "file_server",
                                                    "root": "/tpdata/trojan-panel/webfile/",
                                                    "index_names": [
                                                        "index.html","index.htm"
                                                    ]
                                                }
                                            ],
                                            "terminal": true
                                        }
                                    ]
                                }
                            ]
                        }
                    ],
                    "tls_connection_policies": [
                        {
                            "match": {
                                "sni": [
                                    "${ip}"
                                ]
                            }
                        }
                    ],
                    "automatic_https": {
                        "disable": true
                    }
                }
            }
        },
        "tls": {
            "certificates": {
                "load_files": [
                    {
                        "certificate": "${certificate_path}",
                        "key": "${key_path}"
                    }
                ]
            }
        }
    }
}`
	configContent = strings.ReplaceAll(configContent, "${ip}", strconv.FormatInt(int64(naiveProxyConfigDto.Port), 10))
	configContent = strings.ReplaceAll(configContent, "${crt_path}", certConfig.CrtPath)
	configContent = strings.ReplaceAll(configContent, "${key_path}", certConfig.KeyPath)

	naiveProxyConfig := &bo.NaiveProxyConfig{}
	// 将json字符串映射到模板对象
	if err = json.Unmarshal([]byte(configContent), naiveProxyConfig); err != nil {
		logrus.Errorf("naiveproxy template config反序列化异常 err: %v", err)
		return err
	}

	accountAuthVos, err := dao.SelectAccountUsernameAndPass()
	if err != nil {
		return err
	}
	if len(accountAuthVos) > 0 {
		routeHandles := make([]bo.RouteHandle, 0)
		for _, item := range accountAuthVos {
			handle := bo.RouteHandle{
				AuthUserDeprecated: item.Username,
				AuthPassDeprecated: item.Pass,
				Handler:            bo.TypeMessage("forward_proxy"),
				HideIp:             bo.TypeMessage("true"),
				HideVia:            bo.TypeMessage("true"),
				ProbeResistance:    bo.TypeMessage("{}"),
			}
			routeHandles = append(routeHandles, handle)
		}
		handles := naiveProxyConfig.Apps.Http.Servers.Srv0.Routes[0].Handle[0].HandleRoutes[0].Handle
		handles = append(handles, routeHandles...)
	}

	configContentByte, err := json.MarshalIndent(naiveProxyConfig, "", "    ")
	if err != nil {
		logrus.Errorf("naiveproxy template config反序列化异常 err: %v", err)
		return err
	}

	_, err = file.Write(configContentByte)
	if err != nil {
		logrus.Errorf("naiveproxy config-%d.json文件写入异常 err: %v", naiveProxyConfigDto.ApiPort, err)
		return err
	}
	return nil
}

func InitNaiveProxyBinFile() error {
	// 初始化文件夹
	naiveProxyPath := constant.NaiveProxyPath
	if !util.Exists(naiveProxyPath) {
		if err := os.MkdirAll(naiveProxyPath, os.ModePerm); err != nil {
			logrus.Errorf("创建NaiveProxy文件夹异常 err: %v", err)
			return err
		}
	}

	// 下载二进制文件
	binaryFilePath, err := util.GetBinaryFilePath(constant.NaiveProxy)
	if err != nil {
		return err
	}
	if !util.Exists(binaryFilePath) {
		if err = util.DownloadFile(fmt.Sprintf("%s/naiveproxy-%s-%s", constant.DownloadBaseUrl, runtime.GOOS, runtime.GOARCH),
			binaryFilePath); err != nil {
			logrus.Errorf("NaiveProxy二进制文件下载失败 err: %v", err)
			return err
		}
	}
	return nil
}
