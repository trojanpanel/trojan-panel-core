package xray

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/app/proxyman/command"
	statsService "github.com/xtls/xray-core/app/stats/command"
	statscmd "github.com/xtls/xray-core/app/stats/command"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/proxy/shadowsocks"
	"github.com/xtls/xray-core/proxy/socks"
	"github.com/xtls/xray-core/proxy/trojan"
	"github.com/xtls/xray-core/proxy/vless"
	"github.com/xtls/xray-core/proxy/vmess"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
	"time"
	"trojan-panel-core/model/constant"
	"trojan-panel-core/model/dto"
	"trojan-panel-core/model/vo"
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	clo = func() {
		cancel()
		if conn != nil {
			conn.Close()
		}
	}
	if err != nil {
		logrus.Errorf("xray apiClient init err: %v", err)
		err = errors.New(constant.GrpcError)
	}
	return
}

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
		logrus.Errorf("xray QueryStats err: %v", err)
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

// GetBoundStats query inbound/outbound status
func (x *xrayApi) GetBoundStats(bound string, tag string, link string, reset bool) (*vo.XrayStatsVo, error) {
	conn, ctx, clo, err := apiClient(x.apiPort)
	defer clo()
	if err != nil {
		return nil, err
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
		logrus.Errorf("xray GetBoundStats err: %v", err)
		return nil, errors.New(constant.GrpcError)
	}
	statsVo := vo.XrayStatsVo{
		Name:  tag,
		Value: int(downLinkResponse.GetStat().GetValue()),
	}
	return &statsVo, nil
}

// GetUserStats query account status
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
		logrus.Errorf("xray GetUserStats err: %v", err)
		return nil, errors.New(constant.GrpcError)
	}
	statsVo := vo.XrayStatsVo{
		Name:  email,
		Value: int(downLinkResponse.GetStat().GetValue()),
	}
	return &statsVo, nil
}

// AddUser add user
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

	nodeConfig, err := service.SelectNodeConfigByNodeTypeIdAndApiPort(x.apiPort, constant.Xray)
	if err != nil {
		return nil
	}

	handlerServiceClient := command.NewHandlerServiceClient(conn)
	switch dto.Protocol {
	case constant.ProtocolShadowsocks:
		var cipherType shadowsocks.CipherType
		switch nodeConfig.XraySSMethod {
		case "aes-128-gcm":
			cipherType = shadowsocks.CipherType_AES_128_GCM
		case "aes-256-gcm":
			cipherType = shadowsocks.CipherType_AES_256_GCM
		case "chacha20-poly1305":
			cipherType = shadowsocks.CipherType_CHACHA20_POLY1305
		case "xchacha20-poly1305":
			cipherType = shadowsocks.CipherType_XCHACHA20_POLY1305
		case "none":
			cipherType = shadowsocks.CipherType_NONE
		default:
			cipherType = shadowsocks.CipherType_UNKNOWN
		}

		_, err = handlerServiceClient.AlterInbound(ctx, &command.AlterInboundRequest{
			Tag: "user",
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: dto.Password,
						Level: 0,
						Account: serial.ToTypedMessage(&shadowsocks.Account{
							Password:   dto.Password,
							CipherType: cipherType,
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
							Flow:     nodeConfig.XrayFlow,
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
							Flow:       nodeConfig.XrayFlow,
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
	case constant.ProtocolSocks:
		_, err = handlerServiceClient.AlterInbound(ctx, &command.AlterInboundRequest{
			Tag: "user",
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: dto.Password,
						Level: 0,
						Account: serial.ToTypedMessage(&socks.Account{
							Password: dto.Password,
						}),
					},
				}),
		})
	}
	if err != nil {
		if strings.HasSuffix(err.Error(), "already exists.") {
			return nil
		}
		logrus.Errorf("xray AddUser err: %v", err)
		return errors.New(constant.GrpcError)
	}
	return nil
}

// RemoveInboundHandler delete inbound
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
		logrus.Errorf("xray RemoveInboundHandler err: %v", err)
		return errors.New(constant.GrpcError)
	}
	return nil
}

// DeleteUser delete user
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
		logrus.Errorf("xray DeleteUser err: %v", err)
		return errors.New(constant.GrpcError)
	}
	return nil
}

// GetSysStats get running status
func (x *xrayApi) GetSysStats() (stats *statsService.SysStatsResponse, err error) {
	conn, ctx, clo, err := apiClient(x.apiPort)
	defer clo()
	if err != nil {
		return nil, err
	}
	statsServiceClient := statscmd.NewStatsServiceClient(conn)
	sysStats, err := statsServiceClient.GetSysStats(ctx, &statsService.SysStatsRequest{})
	if err != nil {
		if strings.HasSuffix(err.Error(), "not found.") {
			return nil, nil
		}
		logrus.Errorf("xray GetSysStats err: %v", err)
		return nil, errors.New(constant.GrpcError)
	}
	return sysStats, nil
}
