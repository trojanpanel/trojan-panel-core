package xray

import (
	"context"
	"fmt"
	statscmd "github.com/xtls/xray-core/app/stats/command"
	"github.com/xtls/xray-core/common"
	"github.com/xtls/xray-core/common/net"
	"google.golang.org/grpc"
)

func RemoveHandler() {

}

func AddRemoveUser() {
}
func QueryStats() {

}

func GetStats(ip string, port net.Port, tag string, email string) {
	cmdConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithInsecure(), grpc.WithBlock())
	common.Must(err)
	defer cmdConn.Close()

	name := fmt.Sprintf("%s>>>%s>>>traffic>>>downlink", tag, email)
	sClient := statscmd.NewStatsServiceClient(cmdConn)

	sresp, err := sClient.GetStats(context.Background(), &statscmd.GetStatsRequest{
		Name:   name,
		Reset_: false,
	})
	common.Must(err)
	fmt.Println(sresp.Stat)
}
