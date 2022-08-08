package trojango

import (
	"context"
	"errors"
	"fmt"
	"github.com/p4gefau1t/trojan-go/api/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"trojan-panel-core/module/constant"
)

type trojanGoApi struct {
	ctx     context.Context
	apiPort string
}

// NewTrojanGoApi 初始化Trojan Go Api
func NewTrojanGoApi(apiPort string) *trojanGoApi {
	return &trojanGoApi{
		ctx:     context.Background(),
		apiPort: apiPort,
	}
}

func apiClient(apiPort string) (service.TrojanServerServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%s", apiPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Errorf("Trojan Go gRPC初始化失败 err: %v\n", err)
		return nil, nil, errors.New(constant.GrpcError)
	}
	return service.NewTrojanServerServiceClient(conn), conn, nil
}

// ListUsers 查询节点上的所有用户
func (t *trojanGoApi) ListUsers() ([]*service.UserStatus, error) {
	client, conn, err := apiClient(t.apiPort)
	if err != nil {
		return nil, err
	}
	stream, err := client.ListUsers(t.ctx, &service.ListUsersRequest{})
	if err != nil {
		logrus.Errorf("trojan go list users stream err: %v\n", err)
		return nil, errors.New(constant.GrpcError)
	}
	defer func() {
		stream.CloseSend()
		conn.Close()
	}()

	var userStatus []*service.UserStatus
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			logrus.Errorf("trojan go list users recv err: %v\n", err)
		}
		userStatus = append(userStatus, resp.Status)
	}
	return userStatus, nil
}

// GetUser 查询节点上的用户
func (t *trojanGoApi) GetUser(hash string) (*service.UserStatus, error) {
	client, conn, err := apiClient(t.apiPort)
	if err != nil {
		return nil, err
	}
	stream, err := client.GetUsers(t.ctx)
	if err != nil {
		logrus.Errorf("trojan go get user stream err: %v\n", err)
		return nil, errors.New(constant.GrpcError)
	}
	defer func() {
		stream.CloseSend()
		conn.Close()
	}()

	err = stream.Send(&service.GetUsersRequest{
		User: &service.User{
			Hash: hash,
		},
	})
	if err != nil {
		logrus.Errorf("trojan go get users stream send err: %v\n", err)
		return nil, errors.New(constant.GrpcError)
	}
	resp, err := stream.Recv()
	if err != nil {
		logrus.Errorf("trojan go get users stream send err: %v\n", err)
		return nil, errors.New(constant.GrpcError)
	}
	return resp.Status, nil
}

// 节点上设置用户
func (t *trojanGoApi) setUser(setUsersRequest *service.SetUsersRequest) error {
	client, conn, err := apiClient(t.apiPort)
	if err != nil {
		return err
	}
	stream, err := client.SetUsers(t.ctx)
	if err != nil {
		logrus.Errorf("trojan go set users stream err: %v\n", err)
		return errors.New(constant.GrpcError)
	}
	defer func() {
		stream.CloseSend()
		conn.Close()
	}()

	err = stream.Send(setUsersRequest)
	if err != nil {
		logrus.Errorf("trojan go set user stream send err: %v\n", err)
		return errors.New(constant.GrpcError)
	}
	resp, err := stream.Recv()
	if err != nil {
		logrus.Errorf("trojan go set user stream recv err: %v\n", err)
		return errors.New(constant.GrpcError)
	}
	if !resp.Success {
		logrus.Errorf("trojan go set user fail err: %v\n", err)
		// 重试
		return errors.New(constant.GrpcError)
	}
	return nil
}

// SetUserIpLimit 节点上设置用户设备数
func (t *trojanGoApi) SetUserIpLimit(hash string, ipLimit uint) error {
	req := &service.SetUsersRequest{
		Status: &service.UserStatus{
			User: &service.User{
				Hash: hash,
			},
			IpLimit: int32(ipLimit),
		},
		Operation: service.SetUsersRequest_Modify,
	}
	return t.setUser(req)
}

// SetUserSpeedLimit 节点上设置用户限速
func (t *trojanGoApi) SetUserSpeedLimit(hash string, uploadSpeedLimit int, downloadSpeedLimit int) error {
	req := &service.SetUsersRequest{
		Status: &service.UserStatus{
			User: &service.User{
				Hash: hash,
			},
			SpeedLimit: &service.Speed{
				UploadSpeed:   uint64(uploadSpeedLimit),
				DownloadSpeed: uint64(downloadSpeedLimit),
			},
		},
		Operation: service.SetUsersRequest_Modify,
	}
	return t.setUser(req)
}

// DeleteUser 节点上删除用户
func (t *trojanGoApi) DeleteUser(hash string) error {
	req := &service.SetUsersRequest{
		Status: &service.UserStatus{
			User: &service.User{
				Hash: hash,
			},
		},
		Operation: service.SetUsersRequest_Delete,
	}
	return t.setUser(req)
}

// AddUser 节点上添加用户
func (t *trojanGoApi) AddUser(hash string, ipLimit int, uploadSpeedLimit int, downloadSpeedLimit int) error {
	req := &service.SetUsersRequest{
		Status: &service.UserStatus{
			User: &service.User{
				Hash: hash,
			},
			IpLimit: int32(ipLimit),
			SpeedLimit: &service.Speed{
				UploadSpeed:   uint64(uploadSpeedLimit),
				DownloadSpeed: uint64(downloadSpeedLimit),
			},
		},
		Operation: service.SetUsersRequest_Add,
	}
	return t.setUser(req)
}
