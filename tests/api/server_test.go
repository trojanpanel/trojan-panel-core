package api

import (
	"fmt"
	"os"
	"testing"
	"time"
	"trojan-core/api/server"
	"trojan-core/model/constant"
)

func TestGetServerStats(t *testing.T) {
	token := ""
	conn, ctx, clo, err := newGrpcInstance(token, fmt.Sprintf("127.0.0.1:%s", os.Getenv(constant.GrpcPort)), 4*time.Second)
	defer clo()
	if err != nil {
		fmt.Println(err.Error())
	}
	client := server.NewApiServerServiceClient(conn)
	dto := server.ApiServerDto{}
	send, err := client.GetServerStats(ctx, &dto)
	if err != nil {
		fmt.Println(err.Error())
	}
	if send.Success {
		fmt.Println(send.Data)
	}
}
