package xray

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/app/proxyman/command"
	statscmd "github.com/xtls/xray-core/app/stats/command"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/proxy/shadowsocks"
	"github.com/xtls/xray-core/proxy/trojan"
	"github.com/xtls/xray-core/proxy/vless"
	"github.com/xtls/xray-core/proxy/vmess"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
	"time"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/module/vo"
	"trojan-panel-core/service"
	"trojan-panel-core/util"
)

type xrayApi struct {
	apiPort uint
}

func NewXrayApi(apiPort uint) *xrayApi {
	return &xrayApi{
		apiPort: apiPort,
	}
}

func apiClient(apiPort uint) (conn *grpc.ClientConn, ctx context.Context, clo func(), err error) {
	conn, err = grpc.Dial(fmt.Sprintf("127.0.0.1:%d", apiPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	clo = func() {
		cancel()
		conn.Close()
	}
	if err != nil {
		logrus.Errorf("gRPC初始化失败 err: %v", err)
		err = errors.New(constant.GrpcError)
	}
	return
}

// QueryStats 全量状态
func (x *xrayApi) QueryStats(pattern string, reset bool) ([]vo.XrayStatsVo, error) {
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
		logrus.Errorf("xray query stats err: %v", err)
		return nil, errors.New(constant.GrpcError)
	}

	stats := response.GetStat()
	var xrayStatsVos []vo.XrayStatsVo
	for _, stat := range stats {
		xrayStatsVos = append(xrayStatsVos, vo.XrayStatsVo{
			Name:  stat.Name,
			Value: int(stat.GetValue()),
		})
	}
	return xrayStatsVos, nil
}

// GetBoundStats 查询入/出站状态
func (x *xrayApi) GetBoundStats(bound string, tag string, link string, reset bool) (*vo.XrayStatsVo, error) {
	conn, ctx, clo, err := apiClient(x.apiPort)
	defer clo()
	if err != nil {
		return nil, nil
	}
	statsServiceClient := statscmd.NewStatsServiceClient(conn)
	downLinkResponse, err := statsServiceClient.GetStats(ctx, &statscmd.GetStatsRequest{
		Name:   fmt.Sprintf("%s>>>%s>>>traffic>>>%s", bound, tag, link),
		Reset_: reset,
	})
	if err != nil {
		if strings.HasSuffix(err.Error(), "not found.") {
			return nil, nil
		}
		logrus.Errorf("xray get bound stats err: %v", err)
		return nil, errors.New(constant.GrpcError)
	}
	statsVo := vo.XrayStatsVo{
		Name:  tag,
		Value: int(downLinkResponse.GetStat().GetValue()),
	}
	return &statsVo, nil
}

// GetUserStats 查询用户状态
func (x *xrayApi) GetUserStats(email string, link string, reset bool) (*vo.XrayStatsVo, error) {
	conn, ctx, clo, err := apiClient(x.apiPort)
	defer clo()
	if err != nil {
		return nil, err
	}
	statsServiceClient := statscmd.NewStatsServiceClient(conn)
	downLinkResponse, err := statsServiceClient.GetStats(ctx, &statscmd.GetStatsRequest{
		Name:   fmt.Sprintf("user>>>%s>>>traffic>>>%s", email, link),
		Reset_: reset,
	})
	if err != nil {
		if strings.HasSuffix(err.Error(), "not found.") {
			return nil, nil
		}
		logrus.Errorf("xray get user stats err: %v", err)
		return nil, errors.New(constant.GrpcError)
	}
	statsVo := vo.XrayStatsVo{
		Name:  email,
		Value: int(downLinkResponse.GetStat().GetValue()),
	}
	return &statsVo, nil
}

// AddUser 添加用户
func (x *xrayApi) AddUser(dto dto.XrayAddUserDto) error {
	xrayStatsVo, err := x.GetUserStats(dto.Password, "downlink", false)
	if err != nil {
		return err
	}
	if xrayStatsVo != nil {
		return nil
	}
	conn, ctx, clo, err := apiClient(x.apiPort)
	defer clo()
	if err != nil {
		return nil
	}
	xrayTemplate, err := service.SelectXrayTemplate()
	if err != nil {
		return err
	}
	handlerServiceClient := command.NewHandlerServiceClient(conn)
	switch dto.Protocol {
	case constant.ProtocolShadowsocks:
		_, err = handlerServiceClient.AlterInbound(ctx, &command.AlterInboundRequest{
			Tag: "user",
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: dto.Password,
						Level: 0,
						Account: serial.ToTypedMessage(&shadowsocks.Account{
							Password:   dto.Password,
							CipherType: dto.CipherType,
						}),
					},
				}),
		})
	case constant.ProtocolTrojan:
		_, err = handlerServiceClient.AlterInbound(ctx, &command.AlterInboundRequest{
			Tag: "user",
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: dto.Password,
						Level: 0,
						Account: serial.ToTypedMessage(&trojan.Account{
							Password: dto.Password,
							Flow:     xrayTemplate.Flow,
						}),
					},
				}),
		})
	case constant.ProtocolVless:
		_, err = handlerServiceClient.AlterInbound(ctx, &command.AlterInboundRequest{
			Tag: "user",
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: dto.Password,
						Level: 0,
						Account: serial.ToTypedMessage(&vless.Account{
							Id:         util.GenerateUUID(dto.Password),
							Flow:       xrayTemplate.Flow,
							Encryption: "none",
						}),
					},
				}),
		})
	case constant.ProtocolVmess:
		_, err = handlerServiceClient.AlterInbound(ctx, &command.AlterInboundRequest{
			Tag: "user",
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: dto.Password,
						Level: 0,
						Account: serial.ToTypedMessage(&vmess.Account{
							Id:      util.GenerateUUID(dto.Password),
							AlterId: 0,
						}),
					},
				}),
		})
	}
	if err != nil {
		if strings.HasSuffix(err.Error(), "already exists.") {
			return nil
		}
		logrus.Errorf("xray add user err: %v", err)
		return errors.New(constant.GrpcError)
	}
	return nil
}

// RemoveInboundHandler 删除入站
func (x *xrayApi) RemoveInboundHandler(tag string) error {
	conn, ctx, clo, err := apiClient(x.apiPort)
	defer clo()
	if err != nil {
		return nil
	}
	handlerServiceClient := command.NewHandlerServiceClient(conn)
	_, err = handlerServiceClient.RemoveInbound(ctx, &command.RemoveInboundRequest{
		Tag: tag,
	})
	if err != nil {
		logrus.Errorf("xray remove inbound err: %v", err)
		return errors.New(constant.GrpcError)
	}
	return nil
}

// DeleteUser 删除用户
func (x *xrayApi) DeleteUser(email string) error {
	xrayStatsVo, err := x.GetUserStats(email, "downlink", false)
	if err != nil {
		return err
	}
	if xrayStatsVo == nil {
		return nil
	}
	conn, ctx, clo, err := apiClient(x.apiPort)
	defer clo()
	if err != nil {
		return nil
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
		logrus.Errorf("xray remove user err: %v", err)
		return errors.New(constant.GrpcError)
	}
	return nil
}
