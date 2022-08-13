package xray

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/app/proxyman"
	"github.com/xtls/xray-core/app/proxyman/command"
	statscmd "github.com/xtls/xray-core/app/stats/command"
	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/proxy/shadowsocks"
	"github.com/xtls/xray-core/proxy/trojan"
	"github.com/xtls/xray-core/proxy/vless"
	"github.com/xtls/xray-core/proxy/vless/inbound"
	"github.com/xtls/xray-core/proxy/vmess"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/module/vo"
)

type xrayApi struct {
	apiPort int
}

func NewXrayApi(apiPort int) *xrayApi {
	return &xrayApi{
		apiPort: apiPort,
	}
}

func apiClient(apiPort int) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%s", apiPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Errorf("Xray gRPC初始化失败 err: %v\n", err)
		return nil, errors.New(constant.GrpcError)
	}
	return conn, nil
}

// QueryStats 全量状态
func (x *xrayApi) QueryStats(pattern string, reset bool) ([]vo.XrayStatsVo, error) {
	conn, err := apiClient(x.apiPort)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	statsServiceClient := statscmd.NewStatsServiceClient(conn)
	response, err := statsServiceClient.QueryStats(context.Background(), &statscmd.QueryStatsRequest{
		Pattern: pattern,
		Reset_:  reset,
	})
	if err != nil {
		logrus.Errorf("xray query stats err: %v\n", err)
		return nil, errors.New(constant.GrpcError)
	}

	stats := response.GetStat()
	var xrayStatsVos []vo.XrayStatsVo
	for _, stat := range stats {
		xrayStatsVos = append(xrayStatsVos, vo.XrayStatsVo{
			Name:  stat.Name,
			Value: stat.GetValue(),
		})
	}
	return xrayStatsVos, nil
}

// GetUserStats 查询用户状态
func (x *xrayApi) GetUserStats(email string, link string, reset bool) (*vo.XrayStatsVo, error) {
	conn, err := apiClient(x.apiPort)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	statsServiceClient := statscmd.NewStatsServiceClient(conn)
	downLinkResponse, err := statsServiceClient.GetStats(context.Background(), &statscmd.GetStatsRequest{
		Name:   fmt.Sprintf("user>>>%s>>>traffic>>>%s", email, link),
		Reset_: reset,
	})
	if err != nil {
		logrus.Errorf("xray get user stats err: %v\n", err)
		return nil, errors.New(constant.GrpcError)
	}
	statsVo := vo.XrayStatsVo{
		Name:  email,
		Value: downLinkResponse.GetStat().GetValue(),
	}
	return &statsVo, nil
}

// GetBoundStats 查询入/出站状态
func (x *xrayApi) GetBoundStats(bound string, tag string, link string, reset bool) (*vo.XrayStatsVo, error) {
	conn, err := apiClient(x.apiPort)
	if err != nil {
		return nil, nil
	}
	defer conn.Close()
	statsServiceClient := statscmd.NewStatsServiceClient(conn)
	downLinkResponse, err := statsServiceClient.GetStats(context.Background(), &statscmd.GetStatsRequest{
		Name:   fmt.Sprintf("%s>>>%s>>>traffic>>>%s", bound, tag, link),
		Reset_: reset,
	})
	if err != nil {
		logrus.Errorf("xray get bound stats err: %v\n", err)
		return nil, errors.New(constant.GrpcError)
	}
	statsVo := vo.XrayStatsVo{
		Name:  tag,
		Value: downLinkResponse.GetStat().GetValue(),
	}
	return &statsVo, nil
}

// AddInboundHandler 添加入站
func (x *xrayApi) AddInboundHandler(dto dto.XrayAddBoundDto) error {
	conn, err := apiClient(x.apiPort)
	if err != nil {
		return err
	}
	defer conn.Close()
	handlerServiceClient := command.NewHandlerServiceClient(conn)
	addInboundResponse, err := handlerServiceClient.AddInbound(context.Background(), &command.AddInboundRequest{
		Inbound: &core.InboundHandlerConfig{
			Tag: dto.Tag,
			ReceiverSettings: serial.ToTypedMessage(&proxyman.ReceiverConfig{
				PortList: &net.PortList{Range: []*net.PortRange{net.SinglePortRange(net.Port(dto.Port))}},
				Listen:   net.NewIPOrDomain(net.LocalHostIP),
			}),
			ProxySettings: serial.ToTypedMessage(&inbound.Config{
				Clients: []*protocol.User{
					{
						Level: 0,
						Account: serial.ToTypedMessage(&vmess.Account{
							Id: "0cdf8a45-303d-4fed-9780-29aa7f54175e",
							SecuritySettings: &protocol.SecurityConfig{
								Type: protocol.SecurityType_AES128_GCM,
							},
						}),
					},
				},
			}),
		},
	})
	if err != nil {
		logrus.Errorf("xray add inbound err: %v\n", err)
		return errors.New(constant.GrpcError)
	}
	if addInboundResponse == nil {
		logrus.Errorf("xray add inbound unexpected nil response")
		return errors.New(constant.GrpcError)
	}
	return nil
}

// RemoveInboundHandler 删除入站
func (x *xrayApi) RemoveInboundHandler(tag string) error {
	conn, err := apiClient(x.apiPort)
	if err != nil {
		return nil
	}
	defer conn.Close()
	handlerServiceClient := command.NewHandlerServiceClient(conn)
	removeInboundResponse, err := handlerServiceClient.RemoveInbound(context.Background(), &command.RemoveInboundRequest{
		Tag: tag,
	})
	if err != nil {
		logrus.Errorf("xray remove inbound err: %v\n", err)
		return errors.New(constant.GrpcError)
	}
	if removeInboundResponse == nil {
		logrus.Errorf("xray remove inbound unexpected nil response")
		return errors.New(constant.GrpcError)
	}
	return nil
}

// AddUser 添加用户
func (x *xrayApi) AddUser(dto dto.XrayAddUserDto) error {
	conn, err := apiClient(x.apiPort)
	if err != nil {
		return nil
	}
	defer conn.Close()
	hsClient := command.NewHandlerServiceClient(conn)
	var resp *command.AlterInboundResponse
	switch dto.Tag {
	case constant.ProtocolShadowsocks:
		resp, _ = hsClient.AlterInbound(context.Background(), &command.AlterInboundRequest{
			Tag: dto.Tag,
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: dto.Email,
						Account: serial.ToTypedMessage(&shadowsocks.Account{
							Password: dto.SSPassword,
						}),
					},
				}),
		})
	case constant.ProtocolTrojan:
		resp, _ = hsClient.AlterInbound(context.Background(), &command.AlterInboundRequest{
			Tag: dto.Tag,
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: dto.Email,
						Account: serial.ToTypedMessage(&trojan.Account{
							Password: dto.TrojanPassword,
						}),
					},
				}),
		})
	case constant.ProtocolVless:
		resp, _ = hsClient.AlterInbound(context.Background(), &command.AlterInboundRequest{
			Tag: dto.Tag,
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: dto.Email,
						Account: serial.ToTypedMessage(&vless.Account{
							Id: dto.VId,
						}),
					},
				}),
		})
	case constant.ProtocolVmess:
		resp, _ = hsClient.AlterInbound(context.Background(), &command.AlterInboundRequest{
			Tag: dto.Tag,
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: dto.Email,
						Account: serial.ToTypedMessage(&vmess.Account{
							Id: dto.VId,
						}),
					},
				}),
		})
	}
	if resp == nil {
		logrus.Errorf("xray add user unexpected nil response")
		return errors.New(constant.GrpcError)
	}
	return nil
}

// DeleteUser 删除用户
func (x *xrayApi) DeleteUser(tag string, email string) error {
	conn, err := apiClient(x.apiPort)
	if err != nil {
		return nil
	}
	defer conn.Close()
	hsClient := command.NewHandlerServiceClient(conn)
	resp, err := hsClient.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag:       tag,
		Operation: serial.ToTypedMessage(&command.RemoveUserOperation{Email: email}),
	})
	if err != nil {
		logrus.Errorf("xray remove user err: %v\n", err)
		return errors.New(constant.GrpcError)
	}
	if resp == nil {
		logrus.Errorf("xray remove user nil response")
		return errors.New(constant.GrpcError)
	}
	return nil
}
