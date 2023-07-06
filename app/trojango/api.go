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
	"time"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
)

type trojanGoApi struct {
	apiPort uint
}

// NewTrojanGoApi 初始化Trojan Go Api
func NewTrojanGoApi(apiPort uint) *trojanGoApi {
	return &trojanGoApi{
		apiPort: apiPort,
	}
}

func apiClient(apiPort uint) (clent service.TrojanServerServiceClient, ctx context.Context, clo func(), err error) {
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", apiPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	clent = service.NewTrojanServerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	clo = func() {
		cancel()
		if conn != nil {
			conn.Close()
		}
	}
	if err != nil {
		logrus.Errorf("Trojan Go gRPC初始化失败 err: %v", err)
		err = errors.New(constant.GrpcError)
	}
	return
}

// ListUsers 查询节点上的所有用户
func (t *trojanGoApi) ListUsers() ([]*service.UserStatus, error) {
	client, ctx, clo, err := apiClient(t.apiPort)
	if err != nil {
		return nil, err
	}
	stream, err := client.ListUsers(ctx, &service.ListUsersRequest{})
	if err != nil {
		logrus.Errorf("trojan go list users stream err: %v", err)
		return nil, errors.New(constant.GrpcError)
	}
	defer func() {
		if stream != nil {
			stream.CloseSend()
		}
		clo()
	}()
	var userStatus []*service.UserStatus
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			// 加入重试机制
			logrus.Errorf("trojan go list users recv err: %v", err)
			return nil, err
		}
		if resp != nil {
			userStatus = append(userStatus, resp.Status)
		}
	}
	return userStatus, nil
}

// GetUser 查询节点上的用户
func (t *trojanGoApi) GetUser(password string) (*service.UserStatus, error) {
	client, ctx, clo, err := apiClient(t.apiPort)
	if err != nil {
		return nil, err
	}
	stream, err := client.GetUsers(ctx)
	if err != nil {
		logrus.Errorf("trojan go get user stream err: %v", err)
		return nil, errors.New(constant.GrpcError)
	}
	defer func() {
		if stream != nil {
			stream.CloseSend()
		}
		clo()
	}()
	err = stream.Send(&service.GetUsersRequest{
		User: &service.User{
			Password: password,
		},
	})
	if err != nil {
		logrus.Errorf("trojan go get users stream send err: %v", err)
		return nil, errors.New(constant.GrpcError)
	}
	resp, err := stream.Recv()
	if resp == nil || err != nil {
		logrus.Errorf("trojan go get users stream recv err: %v", err)
		return nil, errors.New(constant.GrpcError)
	}
	return resp.Status, nil
}

// 节点上设置用户
func (t *trojanGoApi) setUser(setUsersRequest *service.SetUsersRequest) error {
	client, ctx, clo, err := apiClient(t.apiPort)
	if err != nil {
		return err
	}
	stream, err := client.SetUsers(ctx)
	if err != nil {
		logrus.Errorf("trojan go set users stream err: %v", err)
		return errors.New(constant.GrpcError)
	}
	defer func() {
		if stream == nil {
			stream.CloseSend()
		}
		clo()
	}()
	err = stream.Send(setUsersRequest)
	if err != nil {
		logrus.Errorf("trojan go set user stream send err: %v", err)
		return errors.New(constant.GrpcError)
	}
	resp, err := stream.Recv()
	if err != nil {
		logrus.Errorf("trojan go set user recv recv err: %v", err)
		return errors.New(constant.GrpcError)
	}
	if resp != nil && !resp.Success {
		logrus.Errorf("trojan go set user fail resp info: %v", resp.Info)
		// 重试
		return errors.New(constant.GrpcError)
	}
	return nil
}

// ReSetUserTrafficByHash 重设用户流量
func (t *trojanGoApi) ReSetUserTrafficByHash(hash string) error {
	req := &service.SetUsersRequest{
		Status: &service.UserStatus{
			User: &service.User{
				Hash: hash,
			},
			TrafficTotal: &service.Traffic{
				DownloadTraffic: 0,
				UploadTraffic:   0,
			},
		},
		Operation: service.SetUsersRequest_Modify,
	}
	return t.setUser(req)
}

// SetUserIpLimit 节点上设置用户设备数
func (t *trojanGoApi) SetUserIpLimit(password string, ipLimit uint) error {
	req := &service.SetUsersRequest{
		Status: &service.UserStatus{
			User: &service.User{
				Password: password,
			},
			IpLimit: int32(ipLimit),
		},
		Operation: service.SetUsersRequest_Modify,
	}
	return t.setUser(req)
}

// SetUserSpeedLimit 节点上设置用户限速
func (t *trojanGoApi) SetUserSpeedLimit(password string, uploadSpeedLimit int, downloadSpeedLimit int) error {
	req := &service.SetUsersRequest{
		Status: &service.UserStatus{
			User: &service.User{
				Password: password,
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
func (t *trojanGoApi) DeleteUser(password string) error {
	userStatus, err := t.GetUser(password)
	if err != nil {
		return err
	}
	if userStatus == nil {
		return nil
	}
	req := &service.SetUsersRequest{
		Status: &service.UserStatus{
			User: &service.User{
				Password: password,
			},
		},
		Operation: service.SetUsersRequest_Delete,
	}
	return t.setUser(req)
}

// AddUser 节点上添加用户
func (t *trojanGoApi) AddUser(dto dto.TrojanGoAddUserDto) error {
	userStatus, err := t.GetUser(dto.Password)
	if err != nil {
		return err
	}
	if userStatus != nil {
		return nil
	}
	req := &service.SetUsersRequest{
		Status: &service.UserStatus{
			User: &service.User{
				Password: dto.Password,
			},
			TrafficTotal: &service.Traffic{
				UploadTraffic:   uint64(dto.UploadTraffic),
				DownloadTraffic: uint64(dto.DownloadTraffic),
			},
			IpLimit: int32(dto.IpLimit),
			SpeedLimit: &service.Speed{
				UploadSpeed:   uint64(dto.UploadSpeedLimit),
				DownloadSpeed: uint64(dto.DownloadSpeedLimit),
			},
		},
		Operation: service.SetUsersRequest_Add,
	}
	return t.setUser(req)
}
