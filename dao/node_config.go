package dao

import (
	"errors"
	"github.com/sirupsen/logrus"
	"trojan-panel-core/model"
	"trojan-panel-core/model/constant"
)

func SelectNodeConfigByNodeTypeIdAndApiPort(apiPortParam uint, nodeTypeIdParam uint) (*model.NodeConfig, error) {
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
	nodeConfig := model.NodeConfig{
		Id:           id,
		ApiPort:      apiPort,
		NodeTypeId:   nodeTypeId,
		Protocol:     protocol,
		XrayFlow:     xrayFlow,
		XraySSMethod: xraySSMethod,
	}
	return &nodeConfig, nil
}

func InsertNodeConfig(nodeConfig model.NodeConfig) error {
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
