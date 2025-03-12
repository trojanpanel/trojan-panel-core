package proxy

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"trojan-core/model/constant"
	"trojan-core/util"
)

func InitProxy() error {
	if !util.Exists(GetXrayBinPath()) {
		if err := DownloadXray(""); err != nil {
			return err
		}
	}
	if !util.Exists(GetHysteriaBinPath()) {
		if err := DownloadHysteria(""); err != nil {
			return err
		}
	}
	if !util.Exists(GetNaiveProxyBinPath()) {
		if err := DownloadNaiveProxy(""); err != nil {
			return err
		}
	}

	xrayConfigFiles, err := util.ListFiles(constant.XrayConfigDir, constant.XrayConfigExt)
	if err != nil {
		return err
	}
	for _, item := range xrayConfigFiles {
		if err = StartProxy(constant.ProtocolXray, util.GetFileNameWithoutExt(item), item); err != nil {
			logrus.Error(err.Error())
		}
	}
	hysteriaConfigFiles, err := util.ListFiles(constant.HysteriaConfigDir, constant.HysteriaConfigExt)
	if err != nil {
		return err
	}
	for _, item := range hysteriaConfigFiles {
		if err = StartProxy(constant.ProtocolHysteria, util.GetFileNameWithoutExt(item), item); err != nil {
			logrus.Error(err.Error())
		}
	}
	naiveProxyConfigFiles, err := util.ListFiles(constant.NaiveProxyConfigDir, constant.NaiveProxyConfigExt)
	if err != nil {
		return err
	}
	for _, item := range naiveProxyConfigFiles {
		if err = StartProxy(constant.ProtocolNaiveProxy, util.GetFileNameWithoutExt(item), item); err != nil {
			logrus.Error(err.Error())
		}
	}
	return nil
}

func StartProxy(proxy, key, configPath string) error {
	if proxy == constant.ProtocolXray {
		return NewXrayInstance(key, configPath).Start()
	} else if proxy == constant.ProtocolHysteria {
		return NewHysteriaInstance(key, configPath).Start()
	} else if proxy == constant.ProtocolNaiveProxy {
		return NewNaiveProxyInstance(key, configPath).Start()
	}
	return fmt.Errorf("proxy not supported")
}

func StopProxy(proxy, key, configPath string) error {
	if proxy == constant.ProtocolXray {
		return NewXrayInstance(key, configPath).Stop()
	} else if proxy == constant.ProtocolHysteria {
		return NewHysteriaInstance(key, configPath).Stop()
	} else if proxy == constant.ProtocolNaiveProxy {
		return NewNaiveProxyInstance(key, configPath).Stop()
	}
	return fmt.Errorf("proxy not supported")
}
