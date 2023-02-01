package service

import (
	"trojan-panel-core/dao"
	"trojan-panel-core/module"
)

func SelectNodeConfigByNodeTypeIdAndApiPort(apiPort uint, nodeTypeId uint) (*module.NodeConfig, error) {
	return dao.SelectNodeConfigByNodeTypeIdAndApiPort(apiPort, nodeTypeId)
}

func InsertNodeConfig(nodeConfig module.NodeConfig) error {
	return dao.InsertNodeConfig(nodeConfig)
}

func DeleteNodeConfigByNodeTypeIdAndApiPort(apiPort uint, nodeTypeId uint) error {
	return dao.DeleteNodeConfigByNodeTypeIdAndApiPort(apiPort, nodeTypeId)
}
