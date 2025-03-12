package proxy

import (
	"fmt"
	"trojan-core/model/constant"
	"trojan-core/util"
)

func InitProxy() error {
	if err := DownloadXray(""); err != nil {
		return err
	}
	if err := DownloadHysteria(""); err != nil {
		return err
	}
	if err := DownloadNaiveProxy(""); err != nil {
		return err
	}
	return nil
}

func StartProxy(proxy string, key string, value []byte) error {
	if proxy == constant.ProtocolXray {
		configPath := fmt.Sprintf("%s%s.json", constant.XrayConfigDir, key)
		if err := util.SaveBytesToFile(value, configPath); err != nil {
			return err
		}
		return NewXrayInstance(key, configPath).Start()
	} else if proxy == constant.ProtocolHysteria {
		configPath := fmt.Sprintf("%s%s.yaml", constant.HysteriaConfigDir, key)
		if err := util.SaveBytesToFile(value, configPath); err != nil {
			return err
		}
		return NewHysteriaInstance(key, configPath).Start()
	} else if proxy == constant.ProtocolNaiveProxy {
		configPath := fmt.Sprintf("%s%s.json", constant.NaiveProxyConfigDir, key)
		if err := util.SaveBytesToFile(value, configPath); err != nil {
			return err
		}
		return NewNaiveProxyInstance(key, configPath).Start()
	}
	return fmt.Errorf("proxy not supported")
}

func StopProxy(proxy string, key string) error {
	if proxy == constant.ProtocolXray {
		return NewXrayInstance(key, fmt.Sprintf("%s%s.json", constant.XrayConfigDir, key)).Stop()
	} else if proxy == constant.ProtocolHysteria {
		return NewHysteriaInstance(key, fmt.Sprintf("%s%s.yaml", constant.HysteriaConfigDir, key)).Stop()
	} else if proxy == constant.ProtocolNaiveProxy {
		return NewNaiveProxyInstance(key, fmt.Sprintf("%s%s.json", constant.NaiveProxyConfigDir, key)).Stop()
	}
	return fmt.Errorf("proxy not supported")
}
