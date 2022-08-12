package app

import (
	"github.com/sirupsen/logrus"
	"trojan-panel-core/app/hysteria"
	"trojan-panel-core/app/trojango"
	"trojan-panel-core/app/xray"
)

func InitApp() {
	if err := xray.InitXrayApp(); err != nil {
		logrus.Errorf("Xray初始化异常 err: %v\n", err)
	}
	if err := trojango.InitTrojanGoApp(); err != nil {
		logrus.Errorf("TrojanGo初始化异常 err: %v\n", err)
	}
	if err := hysteria.InitHysteriaApp(); err != nil {
		logrus.Errorf("Hysteria初始化异常 err: %v\n", err)
	}
}
