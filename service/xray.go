package service

import (
	"encoding/json"
	"errors"
	"fmt"
	redisgo "github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"os"
	"time"
	"trojan-panel-core/dao/redis"
	"trojan-panel-core/module/bo"
	"trojan-panel-core/module/constant"
)

func SelectXrayTemplate() (bo.XrayTemplate, error) {
	var xrayTemplate bo.XrayTemplate
	bytes, err := redis.Client.String.Get("trojan-panel:config:template-xray").Bytes()
	if err != nil && err != redisgo.ErrNil {
		return xrayTemplate, errors.New(constant.SysError)
	}
	if len(bytes) > 0 {
		if err = json.Unmarshal(bytes, &xrayTemplate); err != nil {
			logrus.Errorln(fmt.Sprintf("SelectXrayTemplate XrayTemplate 反序列化失败 err: %v", err))
			return xrayTemplate, errors.New(constant.SysError)
		}
		return xrayTemplate, nil
	} else {
		xrayTemplateContent, err := os.ReadFile(constant.XrayTemplateFilePath)
		if err != nil {
			logrus.Errorln(fmt.Sprintf("读取Xray模板失败 err: %v", err))
			return xrayTemplate, errors.New(constant.SysError)
		}
		xrayTemplateJson, err := json.Marshal(xrayTemplateContent)
		if err != nil {
			logrus.Errorln(fmt.Sprintf("SelectXrayTemplate XrayTemplate 序列化失败 err: %v", err))
			return xrayTemplate, errors.New(constant.SysError)
		}
		redis.Client.String.Set("trojan-panel:config:template-xray", xrayTemplateJson, time.Minute.Milliseconds()*30/1000)
		return xrayTemplate, nil
	}
}

func UpdateXrayTemplate(xrayTemplate string) error {
	if err := os.WriteFile(constant.XrayTemplateFilePath, []byte(xrayTemplate), 0666); err != nil {
		logrus.Errorln(fmt.Sprintf("写入Xray默认模板异常err: %v", err))
		return errors.New(constant.SysError)
	}
	redis.Client.Key.Del("trojan-panel:config:template-xray")
	return nil
}
