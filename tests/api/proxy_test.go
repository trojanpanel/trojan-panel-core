package api

import (
	"fmt"
	"os"
	"testing"
	"time"
	"trojan-core/api/proxy"
	"trojan-core/model/constant"
)

func TestStartProxy(t *testing.T) {
	token := ""
	conn, ctx, clo, err := newGrpcInstance(token, fmt.Sprintf("127.0.0.1:%s", os.Getenv(constant.GrpcPort)), 4*time.Second)
	defer clo()
	if err != nil {
		fmt.Println(err.Error())
	}
	client := proxy.NewApiProxyServiceClient(conn)
	dto := proxy.StartProxyDto{}
	send, err := client.StartProxy(ctx, &dto)
	if err != nil {
		fmt.Println(err.Error())
	}
	if send.Success {
		fmt.Println(send.Data)
	}
}

func TestStopProxy(t *testing.T) {
	token := ""
	conn, ctx, clo, err := newGrpcInstance(token, fmt.Sprintf("127.0.0.1:%s", os.Getenv(constant.GrpcPort)), 4*time.Second)
	defer clo()
	if err != nil {
		fmt.Println(err.Error())
	}
	client := proxy.NewApiProxyServiceClient(conn)
	dto := proxy.StopProxyDto{}
	send, err := client.StopProxy(ctx, &dto)
	if err != nil {
		fmt.Println(err.Error())
	}
	if send.Success {
		fmt.Println(send.Data)
	}
}
