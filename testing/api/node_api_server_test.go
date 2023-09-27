package api

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	"time"
	"trojan-panel-core/api"
	"trojan-panel-core/model/constant"
)

var (
	conn *grpc.ClientConn
	ctx  context.Context
	clo  func()
	err  error
)

func init() {
	conn, ctx, clo, err = newGrpcInstance("127.0.0.1", 8100, 4*time.Second)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
}

func newGrpcInstance(ip string, grpcPort uint, timeout time.Duration) (conn *grpc.ClientConn, ctx context.Context, clo func(), err error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err = grpc.Dial(fmt.Sprintf("%s:%d", ip, grpcPort), opts...)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	clo = func() {
		cancel()
		if conn != nil {
			conn.Close()
		}
	}
	if err != nil {
		fmt.Printf("gRPC instance init err ip: %s grpcPort: %d err: %v", ip, grpcPort, err)
		err = errors.New(constant.GrpcError)
	}
	return
}

func TestAddNode(t *testing.T) {
	defer clo()

	nodeAddDto := api.NodeAddDto{
		Port:                  8089,
		Domain:                "demo.ioerror.top",
		NodeTypeId:            constant.Hysteria2,
		Hysteria2ObfsPassword: "123456",
		Hysteria2UpMbps:       100,
		Hysteria2DownMbps:     100,
	}
	client := api.NewApiNodeServiceClient(conn)
	resp, err := client.AddNode(ctx, &nodeAddDto)
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
	defer clo()
	nodeRemoveDto := api.NodeRemoveDto{
		NodeTypeId: constant.Hysteria2,
		Port:       8089,
	}
	client := api.NewApiNodeServiceClient(conn)
	resp, err := client.RemoveNode(ctx, &nodeRemoveDto)
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
