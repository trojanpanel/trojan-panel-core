package hysteria

import (
	"github.com/sirupsen/logrus"
	"os"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

func StartHysteria() {

}

// 初始化Hysteria
func init() {
	hysteriaPath := constant.HysteriaPath
	if !util.Exists(hysteriaPath) {
		if err := os.MkdirAll(hysteriaPath, os.ModePerm); err != nil {
			logrus.Errorf("创建Hysteria文件夹异常 err: %v\n", err)
			panic(err)
		}
	}
}
