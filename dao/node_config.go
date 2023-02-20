package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	redisgo "github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"time"
	"trojan-panel-core/dao/redis"
	"trojan-panel-core/module"
	"trojan-panel-core/module/constant"
)

func SelectNodeConfigByNodeTypeIdAndApiPort(apiPortParam uint, nodeTypeIdParam uint) (*module.NodeConfig, error) {
	stmt, err := sqliteDb.Prepare("select id,node_type_id,api_port,protocol,xray_flow,xray_ss_method from node_config where api_port = ? and node_type_id = ?")
	if err != nil {
		return nil, errors.New(constant.SysError)
	}
	rows, err := stmt.Query(apiPortParam, nodeTypeIdParam)
	if err != nil {
		logrus.Errorf("SelectNodeConfigByNodeTypeIdAndApiPort err: %v", err)
		return nil, errors.New(constant.SysError)
	} else if rows.Err() != nil {
		logrus.Errorf("SelectNodeConfigByNodeTypeIdAndApiPort err: %v", rows.Err())
		return nil, errors.New(constant.SysError)
	}
	defer func() {
		rows.Close()
		stmt.Close()
	}()

	var (
		id           uint
		apiPort      uint
		nodeTypeId   uint
		protocol     string
		xrayFlow     string
		xraySSMethod string
	)
	for rows.Next() {
		if err := rows.Scan(&id, &apiPort, &nodeTypeId, &protocol, &xrayFlow, &xraySSMethod); err != nil {
			return nil, errors.New(constant.SysError)
		}
		break
	}
	nodeConfig := module.NodeConfig{
		Id:           id,
		ApiPort:      apiPort,
		NodeTypeId:   nodeTypeId,
		Protocol:     protocol,
		XrayFlow:     xrayFlow,
		XraySSMethod: xraySSMethod,
	}
	return &nodeConfig, nil
}

func SelectNodeConfig(apiPort uint, nodeTypeId uint) (*module.NodeConfig, error) {
	bytes, err := redis.Client.String.Get(fmt.Sprintf("trojan-panel-core:node-config:%d-%d", apiPort, nodeTypeId)).Bytes()
	if err != nil && err != redisgo.ErrNil {
		return nil, errors.New(constant.SysError)
	}
	if len(bytes) > 0 {
		var nodeConfig module.NodeConfig
		if err = json.Unmarshal(bytes, &nodeConfig); err != nil {
			logrus.Errorln(fmt.Sprintf("SelectNodeConfigByNodeTypeIdAndApiPort NodeConfig 反序列化失败 err: %v", err))
			return nil, errors.New(constant.SysError)
		}
		return &nodeConfig, nil
	} else {
		nodeConfig, err := SelectNodeConfigByNodeTypeIdAndApiPort(apiPort, nodeTypeId)
		if err != nil {
			return nil, err
		}
		nodeConfigJson, err := json.Marshal(*nodeConfig)
		if err != nil {
			logrus.Errorln(fmt.Sprintf("SelectNodeConfigByNodeTypeIdAndApiPort NodeConfig 序列化失败 err: %v", err))
			return nil, errors.New(constant.SysError)
		}
		redis.Client.String.Set(fmt.Sprintf("trojan-panel-core:node-config:%d-%d", apiPort, nodeTypeId), nodeConfigJson, time.Hour.Milliseconds()*48/1000)
		return nodeConfig, nil
	}
}

func InsertNodeConfig(nodeConfig module.NodeConfig) error {
	stmt, err := sqliteDb.Prepare("insert into node_config(node_type_id,api_port,protocol,xray_flow,xray_ss_method) values(?,?,?,?,?)")
	if err != nil {
		return errors.New(constant.SysError)
	}
	defer stmt.Close()
	_, err = stmt.Exec(nodeConfig.NodeTypeId, nodeConfig.ApiPort, nodeConfig.Protocol, nodeConfig.XrayFlow, nodeConfig.XraySSMethod)
	if err != nil {
		return errors.New(constant.SysError)
	}
	return nil
}

func DeleteNodeConfigByNodeTypeIdAndApiPort(apiPort uint, nodeTypeId uint) error {
	stmt, err := sqliteDb.Prepare("delete from node_config where api_port = ? and node_type_id = ?")
	if err != nil {
		return errors.New(constant.SysError)
	}
	defer stmt.Close()
	_, err = stmt.Exec(apiPort, nodeTypeId)
	if err != nil {
		return errors.New(constant.SysError)
	}
	return nil
}
