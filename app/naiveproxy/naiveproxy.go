package naiveproxy

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"trojan-panel-core/core"
	"trojan-panel-core/core/process"
	"trojan-panel-core/model/constant"
	"trojan-panel-core/model/dto"
	"trojan-panel-core/util"
)

func InitNaiveProxyApp() error {
	apiPorts, err := util.GetConfigApiPorts(constant.NaiveProxyPath)
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

func StopNaiveProxy(apiPort uint, removeFile bool) error {
	if err := process.NewNaiveProxyInstance().Stop(apiPort, removeFile); err != nil {
		logrus.Errorf("naiveproxy stop err: %v", err)
		return err
	}
	return nil
}

func RestartNaiveProxy(apiPort uint) error {
	if err := StopNaiveProxy(apiPort, false); err != nil {
		return err
	}
	if err := StartNaiveProxy(dto.NaiveProxyConfigDto{ApiPort: apiPort}); err != nil {
		return err
	}
	return nil
}

func initNaiveProxy(naiveProxyConfigDto dto.NaiveProxyConfigDto) error {
	naiveProxyConfigFilePath, err := util.GetConfigFilePath(constant.NaiveProxy, naiveProxyConfigDto.ApiPort)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(naiveProxyConfigFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		logrus.Errorf("create naiveproxy %s file err: %v", naiveProxyConfigFilePath, err)
		return err
	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	certConfig := core.Config.CertConfig
	configContent := `{
    "admin": {
        "disabled": false,
		"listen": "127.0.0.1:${api_port}"
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
											"handle": []
										},
                                        {
                                            "match": [
                                                {
                                                    "host": [
                                                        "${domain}"
                                                    ]
                                                }
                                            ],
                                            "handle": [
                                                {
                                                    "handler": "file_server",
                                                    "root": "/tpdata/web/",
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
                                    "${domain}"
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
                        "certificate": "${crt_path}",
                        "key": "${key_path}"
                    }
                ]
            }
        }
    }
}`
	configContent = strings.ReplaceAll(configContent, "${api_port}", strconv.FormatInt(int64(naiveProxyConfigDto.ApiPort), 10))
	configContent = strings.ReplaceAll(configContent, "${domain}", naiveProxyConfigDto.Domain)
	configContent = strings.ReplaceAll(configContent, "${port}", strconv.FormatInt(int64(naiveProxyConfigDto.Port), 10))
	configContent = strings.ReplaceAll(configContent, "${crt_path}", certConfig.CrtPath)
	configContent = strings.ReplaceAll(configContent, "${key_path}", certConfig.KeyPath)

	_, err = file.WriteString(configContent)
	if err != nil {
		logrus.Errorf("naiveproxy file config-%d.json write err: %v", naiveProxyConfigDto.ApiPort, err)
		return err
	}
	return nil
}

func InitNaiveProxyBinFile() error {
	naiveProxyPath := constant.NaiveProxyPath
	if !util.Exists(naiveProxyPath) {
		if err := os.MkdirAll(naiveProxyPath, os.ModePerm); err != nil {
			logrus.Errorf("create navieproxy folder err: %v", err)
			return err
		}
	}

	binaryFilePath, err := util.GetBinaryFilePath(constant.NaiveProxy)
	if err != nil {
		return err
	}
	if !util.Exists(binaryFilePath) {
		logrus.Errorf("naiveproxy binary file does not exist")
		return errors.New(constant.BinaryFileNotExist)
	}
	return nil
}
