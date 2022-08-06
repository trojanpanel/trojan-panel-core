package trojango

import (
	"github.com/sirupsen/logrus"
	"os"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/util"
)

// 初始化TrojanGo
func init() {
	trojanGoPath := constant.TrojanGoPath
	if !util.Exists(trojanGoPath) {
		if err := os.MkdirAll(trojanGoPath, os.ModePerm); err != nil {
			logrus.Errorf("创建trojango文件夹异常 err: %v\n", err)
			panic(err)
		}
	}

}
