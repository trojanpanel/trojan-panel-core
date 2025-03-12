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

	xrayConfigFiles, err := util.ListFileNames(constant.XrayConfigDir, constant.XrayConfigExt)
	if err != nil {
		return err
	}
	for _, item := range xrayConfigFiles {
		if err = StartProxy(constant.ProtocolXray, item); err != nil {
			logrus.Error(err.Error())
		}
	}
	hysteriaConfigFiles, err := util.ListFileNames(constant.HysteriaConfigDir, constant.HysteriaConfigExt)
	if err != nil {
		return err
	}
	for _, item := range hysteriaConfigFiles {
		if err = StartProxy(constant.ProtocolHysteria, item); err != nil {
			logrus.Error(err.Error())
		}
	}
	naiveProxyConfigFiles, err := util.ListFileNames(constant.NaiveProxyConfigDir, constant.NaiveProxyConfigExt)
	if err != nil {
		return err
	}
	for _, item := range naiveProxyConfigFiles {
		if err = StartProxy(constant.ProtocolNaiveProxy, item); err != nil {
			logrus.Error(err.Error())
		}
	}
	return nil
}

func PrepareConfigFile(proxy, key string, data []byte) error {
	if proxy == constant.ProtocolXray {
		return util.SaveBytesToFile(data, GetXrayConfigPath(key))
	} else if proxy == constant.ProtocolHysteria {
		return util.SaveBytesToFile(data, GetHysteriaConfigPath(key))
	} else if proxy == constant.ProtocolNaiveProxy {
		return util.SaveBytesToFile(data, GetNaiveProxyConfigPath(key))
	}
	return fmt.Errorf("proxy not supported")
}

func StartProxy(proxy, key string) error {
	if proxy == constant.ProtocolXray {
		return NewXrayInstance(key).Start()
	} else if proxy == constant.ProtocolHysteria {
		return NewHysteriaInstance(key).Start()
	} else if proxy == constant.ProtocolNaiveProxy {
		return NewNaiveProxyInstance(key).Start()
	}
	return fmt.Errorf("proxy not supported")
}

func StopProxy(proxy, key string) error {
	if proxy == constant.ProtocolXray {
		return NewXrayInstance(key).Stop()
	} else if proxy == constant.ProtocolHysteria {
		return NewHysteriaInstance(key).Stop()
	} else if proxy == constant.ProtocolNaiveProxy {
		return NewNaiveProxyInstance(key).Stop()
	}
	return fmt.Errorf("proxy not supported")
}
