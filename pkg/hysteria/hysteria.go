package hysteria

import (
	"github.com/sirupsen/logrus"
	"os"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

func StartHysteria() error {
	return nil
}

func StopHysteria() {

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

	hysteriaConfigFilePath := constant.HysteriaConfigFilePath
	if !util.Exists(hysteriaConfigFilePath) {
		file, err := os.Create(hysteriaConfigFilePath)
		if err != nil {
			logrus.Errorf("创建hysteria config.json文件异常 err: %v\n", err)
			panic(err)
		}
		defer file.Close()

		configContent := ``
		_, err = file.WriteString(configContent)
		if err != nil {
			logrus.Errorf("hysteria config.json文件写入异常 err: %v\n", err)
			panic(err)
		}
	}
}
