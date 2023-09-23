package service

import (
	"encoding/json"
	"errors"
	"fmt"
	redisgo "github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
	"trojan-panel-core/dao"
	"trojan-panel-core/dao/redis"
	"trojan-panel-core/model"
	"trojan-panel-core/model/constant"
)

func SelectNodeConfigByNodeTypeIdAndApiPort(apiPort uint, nodeTypeId uint) (*model.NodeConfig, error) {
	bytes, err := redis.Client.String.Get(fmt.Sprintf("trojan-panel-core:node-config:%d-%d", apiPort, nodeTypeId)).Bytes()
	if err != nil && err != redisgo.ErrNil {
		return nil, errors.New(constant.SysError)
	}
	if len(bytes) > 0 {
		var nodeConfig model.NodeConfig
		if err = json.Unmarshal(bytes, &nodeConfig); err != nil {
			logrus.Errorln(fmt.Sprintf("SelectNodeConfigByNodeTypeIdAndApiPort NodeConfig deserialization err: %v", err))
			return nil, errors.New(constant.SysError)
		}
		return &nodeConfig, nil
	} else {
		nodeConfig, err := dao.SelectNodeConfigByNodeTypeIdAndApiPort(apiPort, nodeTypeId)
		if err != nil {
			return nil, err
		}
		nodeConfigJson, err := json.Marshal(*nodeConfig)
		if err != nil {
			logrus.Errorln(fmt.Sprintf("SelectNodeConfigByNodeTypeIdAndApiPort NodeConfig serialization err: %v", err))
			return nil, errors.New(constant.SysError)
		}
		redis.Client.String.Set(fmt.Sprintf("trojan-panel-core:node-config:%d-%d", apiPort, nodeTypeId), nodeConfigJson, time.Hour.Milliseconds()*48/1000)
		return nodeConfig, nil
	}
}

func InsertNodeConfig(nodeConfig model.NodeConfig) error {
	return dao.InsertNodeConfig(nodeConfig)
}

func DeleteNodeConfigByNodeTypeIdAndApiPort(apiPort uint, nodeTypeId uint) error {
	var mutex sync.Mutex
	defer mutex.Unlock()
	if mutex.TryLock() {
		if err := dao.DeleteNodeConfigByNodeTypeIdAndApiPort(apiPort, nodeTypeId); err != nil {
			return err
		}
		if err := redis.Client.Key.RetryDel(fmt.Sprintf("trojan-panel-core:node-config:%d-%d", apiPort, nodeTypeId)); err != nil {
			return err
		}
	}
	return nil
}
