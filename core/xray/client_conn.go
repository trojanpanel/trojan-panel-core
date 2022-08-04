package xray

import (
	"fmt"
	"google.golang.org/grpc"
	"xray-manage/module/constant"
)

var ClientConn *grpc.ClientConn

func InitGrpcClientConn() {
	ClientConn, _ = grpc.Dial(fmt.Sprintf("127.0.0.1:%s", constant.GrpcPort))
}

func CloseClientConn() {
	if ClientConn != nil {
		ClientConn.Close()
	}
}
