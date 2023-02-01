package dao

import (
	"errors"
	"trojan-panel-core/module"
	"trojan-panel-core/module/constant"
)

func SelectNodeConfigByNodeTypeIdAndApiPort(apiPort uint, nodeTypeId uint) (*module.NodeConfig, error) {
	stmt, err := sqliteDb.Prepare("select id,node_type_id,api_port,protocol,xray_flow,xray_ss_method from node_config where api_port = ? and node_type_id = ?")
	if err != nil {
		return nil, errors.New(constant.SysError)
	}
	row := stmt.QueryRow(apiPort, nodeTypeId)
	if row.Err() != nil {
		return nil, errors.New(constant.SysError)
	}
	var nodeConfig module.NodeConfig
	if err := row.Scan(&nodeConfig); err != nil {
		return nil, errors.New(constant.SysError)
	}
	return &nodeConfig, nil
}

func InsertNodeConfig(nodeConfig module.NodeConfig) error {
	stmt, err := sqliteDb.Prepare("insert into node_config(node_type_id,api_port,protocol,xray_flow,xray_ss_method) values(?,?,?,?,?,?)")
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
