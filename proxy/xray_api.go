package proxy

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/app/proxyman/command"
	statscmd "github.com/xtls/xray-core/app/stats/command"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
	"time"
	"trojan-core/model/bo"
	"trojan-core/model/constant"
)

type XrayApi struct {
	apiPort string
}

func NewXrayApi(apiPort string) *XrayApi {
	return &XrayApi{
		apiPort: apiPort,
	}
}

func apiClient(apiPort string) (conn *grpc.ClientConn, ctx context.Context, clo func(), err error) {
	conn, err = grpc.Dial(fmt.Sprintf("127.0.0.1:%s", apiPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	clo = func() {
		cancel()
		if conn != nil {
			_ = conn.Close()
		}
	}
	if err != nil {
		logrus.Errorf("xray apiClient init err: %v", err)
		err = fmt.Errorf(constant.GrpcError)
	}
	return
}

func (x *XrayApi) QueryStats(pattern string, reset bool) ([]bo.XrayStatsBo, error) {
	conn, ctx, clo, err := apiClient(x.apiPort)
	defer clo()
	if err != nil {
		return nil, err
	}
	statsServiceClient := statscmd.NewStatsServiceClient(conn)
	response, err := statsServiceClient.QueryStats(ctx, &statscmd.QueryStatsRequest{
		Pattern: pattern,
		Reset_:  reset,
	})
	if err != nil {
		logrus.Errorf("xray QueryStats err: %v", err)
		return nil, fmt.Errorf(constant.GrpcError)
	}
	stats := response.GetStat()
	var xrayStatsVos []bo.XrayStatsBo
	for _, stat := range stats {
		xrayStatsVos = append(xrayStatsVos, bo.XrayStatsBo{
			Name:  stat.Name,
			Value: int(stat.GetValue()),
		})
	}
	return xrayStatsVos, nil
}

func (x *XrayApi) AddUser(password string, message *serial.TypedMessage) error {
	conn, ctx, clo, err := apiClient(x.apiPort)
	defer clo()
	if err != nil {
		return nil
	}
	handlerServiceClient := command.NewHandlerServiceClient(conn)
	_, err = handlerServiceClient.AlterInbound(ctx, &command.AlterInboundRequest{
		Tag: "user",
		Operation: serial.ToTypedMessage(&command.AddUserOperation{
			User: &protocol.User{
				Email:   password,
				Level:   0,
				Account: serial.ToTypedMessage(message),
			},
		}),
	})
	if err != nil {
		if strings.HasSuffix(err.Error(), "already exists.") {
			return nil
		}
		logrus.Errorf("xray AddUser err: %v", err)
		return fmt.Errorf(constant.GrpcError)
	}
	return nil
}

func (x *XrayApi) DeleteUser(email string) error {
	conn, ctx, clo, err := apiClient(x.apiPort)
	defer clo()
	if err != nil {
		return err
	}
	hsClient := command.NewHandlerServiceClient(conn)
	_, err = hsClient.AlterInbound(ctx, &command.AlterInboundRequest{
		Tag:       "user",
		Operation: serial.ToTypedMessage(&command.RemoveUserOperation{Email: email}),
	})
	if err != nil {
		if strings.HasSuffix(err.Error(), "not found.") {
			return nil
		}
		logrus.Errorf("xray DeleteUser err: %v", err)
		return fmt.Errorf(constant.GrpcError)
	}
	return nil
}
