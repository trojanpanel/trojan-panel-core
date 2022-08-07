package xray

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/core"
	"trojan-panel-core/module/constant"
)

var server *core.Instance

// StartXray 启动Xray
func StartXray(config *core.Config) error {
	cfgBytes, err := proto.Marshal(config)
	if err != nil {
		logrus.Errorf("解析Xray配置异常 err: %v\n", err)
		return errors.New(constant.XrayConfigError)
	}
	server, err = core.StartInstance("protobuf", cfgBytes)
	if err != nil {
		logrus.Errorf("Xray启动异常 err: %v\n", err)
		return errors.New(constant.XrayStartError)
	}
	return nil
}

// StopXray 关闭Xray
func StopXray() {
	if server != nil {
		server.Close()
	}
}
