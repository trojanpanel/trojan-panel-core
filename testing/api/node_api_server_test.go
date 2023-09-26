package api

import (
	"context"
	"fmt"
	"testing"
	"time"
	"trojan-panel-core/api"
	"trojan-panel-core/model/constant"
)

func TestAddNode(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var server api.NodeApiServer
	nodeAddDto := api.NodeAddDto{
		NodeTypeId:            constant.Hysteria2,
		Hysteria2ObfsPassword: "123456",
		Hysteria2UpMbps:       100,
		Hysteria2DownMbps:     100,
	}
	resp, err := server.AddNode(ctx, &nodeAddDto)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	if resp.GetSuccess() {
		fmt.Printf("result: %v", resp.GetData())
		return
	}
	fmt.Printf("err msg: %s", resp.GetMsg())
}

func TestRemoveNode(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var server api.NodeApiServer
	nodeRemoveDto := api.NodeRemoveDto{
		NodeTypeId: constant.Hysteria2,
		Port:       8089,
	}
	resp, err := server.RemoveNode(ctx, &nodeRemoveDto)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	if resp.GetSuccess() {
		fmt.Printf("result: %v", resp.GetData())
		return
	}
	fmt.Printf("err msg: %s", resp.GetMsg())
}
