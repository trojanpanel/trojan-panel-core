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
	"trojan-core/model/constant"
	"trojan-core/model/dto"
)

type trojanGoApi struct {
	apiPort uint
}

func NewTrojanGoApi(apiPort uint) *trojanGoApi {
	return &trojanGoApi{
		apiPort: apiPort,
	}
}

func apiClient(apiPort uint) (clent service.TrojanServerServiceClient, ctx context.Context, clo func(), err error) {
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", apiPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	clent = service.NewTrojanServerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	clo = func() {
		cancel()
		if conn != nil {
			conn.Close()
		}
	}
	if err != nil {
		logrus.Errorf("trojango apiClient init err: %v", err)
		err = errors.New(constant.GrpcError)
	}
	return
}

// ListUsers query all users on a node
func (t *trojanGoApi) ListUsers() ([]*service.UserStatus, error) {
	client, ctx, clo, err := apiClient(t.apiPort)
	if err != nil {
		return nil, err
	}
	stream, err := client.ListUsers(ctx, &service.ListUsersRequest{})
	if err != nil {
		logrus.Errorf("trojango ListUsers err: %v", err)
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
			logrus.Errorf("trojango ListUsers recv err: %v", err)
			return nil, errors.New(constant.GrpcError)
		}
		if resp != nil {
			userStatus = append(userStatus, resp.Status)
		}
	}
	return userStatus, nil
}

// GetUser query users on a node
func (t *trojanGoApi) GetUser(hash string) (*service.UserStatus, error) {
	client, ctx, clo, err := apiClient(t.apiPort)
	if err != nil {
		return nil, err
	}
	stream, err := client.GetUsers(ctx)
	if err != nil {
		logrus.Errorf("trojango GetUser err: %v", err)
		return nil, errors.New(constant.GrpcError)
	}
	defer func() {
		if stream != nil {
			stream.CloseSend()
		}
		clo()
	}()
	if err = stream.Send(&service.GetUsersRequest{
		User: &service.User{
			Hash: hash,
		},
	}); err != nil {
		logrus.Errorf("trojango GetUser stream send err: %v", err)
		return nil, errors.New(constant.GrpcError)
	}
	resp, err := stream.Recv()
	if resp == nil || err != nil {
		logrus.Errorf("trojango GetUser stream recv err: %v", err)
		return nil, errors.New(constant.GrpcError)
	}
	return resp.Status, nil
}

// set user on node
func (t *trojanGoApi) setUser(setUsersRequest *service.SetUsersRequest) error {
	client, ctx, clo, err := apiClient(t.apiPort)
	if err != nil {
		return err
	}
	stream, err := client.SetUsers(ctx)
	if err != nil {
		logrus.Errorf("trojango setUser err: %v", err)
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
		logrus.Errorf("trojango setUser send err: %v", err)
		return errors.New(constant.GrpcError)
	}
	resp, err := stream.Recv()
	if err != nil {
		logrus.Errorf("trojango setUser recv err: %v", err)
		return errors.New(constant.GrpcError)
	}
	if resp != nil && !resp.Success {
		logrus.Errorf("trojango setUser err resp info: %v", resp.Info)
		return errors.New(constant.GrpcError)
	}
	return nil
}

// ReSetUserTrafficByHash reset user traffic
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

// SetUserIpLimit set the number of user devices on the node
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

// SetUserSpeedLimit set user speed limit on the node
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

// DeleteUser delete user on node
func (t *trojanGoApi) DeleteUser(hash string) error {
	userStatus, err := t.GetUser(hash)
	if err != nil {
		return err
	}
	if userStatus == nil {
		return nil
	}
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

// AddUser add user on node
func (t *trojanGoApi) AddUser(dto dto.TrojanGoAddUserDto) error {
	userStatus, err := t.GetUser(dto.Hash)
	if err != nil {
		return err
	}
	if userStatus != nil {
		return nil
	}
	req := &service.SetUsersRequest{
		Status: &service.UserStatus{
			User: &service.User{
				Hash: dto.Hash,
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
