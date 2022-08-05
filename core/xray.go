package core

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
	"xray-manage/module/constant"
	"xray-manage/module/dto"
	"xray-manage/module/vo"
)

var ClientConn *grpc.ClientConn

// InitGrpcClientConn 初始化gRPC
func InitGrpcClientConn() {
	ClientConn, _ = grpc.Dial(fmt.Sprintf("127.0.0.1:%s", constant.GrpcPort))
}

// CloseClientConn 关闭gRPC
func CloseClientConn() {
	if ClientConn != nil {
		ClientConn.Close()
	}
}

// QueryStats 全量查询状态
func QueryStats(pattern string, reset bool) []vo.XrayStatsVo {
	statsServiceClient := statscmd.NewStatsServiceClient(ClientConn)
	response, err := statsServiceClient.QueryStats(context.Background(), &statscmd.QueryStatsRequest{
		Pattern: pattern,
		Reset_:  reset,
	})
	if err != nil {
		logrus.Errorf("QueryStats获取Stats异常 err: %v\n", err)
		return nil
	}

	stats := response.GetStat()
	var xrayStatsVos []vo.XrayStatsVo
	for _, stat := range stats {
		xrayStatsVos = append(xrayStatsVos, vo.XrayStatsVo{
			Name:  stat.Name,
			Value: stat.GetValue(),
		})
	}
	return xrayStatsVos
}

// GetUserStats 获取用户状态
func GetUserStats(email string, reset bool, link string) *vo.XrayStatsVo {
	statsServiceClient := statscmd.NewStatsServiceClient(ClientConn)
	downLinkResponse, err := statsServiceClient.GetStats(context.Background(), &statscmd.GetStatsRequest{
		Name:   fmt.Sprintf("user>>>%s>>>traffic>>>%s", email, link),
		Reset_: reset,
	})
	if err != nil {
		logrus.Errorf("GetUserStats获取Stats异常 err: %v\n", err)
		return nil
	}
	statsVo := vo.XrayStatsVo{
		Name:  email,
		Value: downLinkResponse.GetStat().GetValue(),
	}
	return &statsVo
}

// GetBoundStats 获取入站状态
func GetBoundStats(bound string, tag string, link string, reset bool) *vo.XrayStatsVo {
	statsServiceClient := statscmd.NewStatsServiceClient(ClientConn)
	downLinkResponse, err := statsServiceClient.GetStats(context.Background(), &statscmd.GetStatsRequest{
		Name:   fmt.Sprintf("%s>>>%s>>>traffic>>>%s", bound, tag, link),
		Reset_: reset,
	})
	if err != nil {
		logrus.Errorf("GetUserStats获取Stats异常 err: %v\n", err)
		return nil
	}
	statsVo := vo.XrayStatsVo{
		Name:  tag,
		Value: downLinkResponse.GetStat().GetValue(),
	}
	return &statsVo
}

// RemoveInboundHandler 删除入站
func RemoveInboundHandler(tag string) error {
	handlerServiceClient := command.NewHandlerServiceClient(ClientConn)
	removeInboundResponse, err := handlerServiceClient.RemoveInbound(context.Background(), &command.RemoveInboundRequest{
		Tag: tag,
	})
	if err != nil {
		logrus.Errorf("删除入站异常 err: %v\n", err)
		return errors.New(constant.XrayRemoveInBoundError)
	}
	if removeInboundResponse == nil {
		logrus.Errorf("unexpected nil response")
		return errors.New(constant.XrayRemoveInBoundError)
	}
	return nil
}

// RemoveOutboundHandler 删除出站
func RemoveOutboundHandler(tag string) error {
	handlerServiceClient := command.NewHandlerServiceClient(ClientConn)
	removeOutboundResponse, err := handlerServiceClient.RemoveOutbound(context.Background(), &command.RemoveOutboundRequest{
		Tag: tag,
	})
	if err != nil {
		logrus.Errorf("删除出站异常 err: %v\n", err)
		return errors.New(constant.XrayRemoveOutBoundError)
	}
	if removeOutboundResponse == nil {
		logrus.Errorf("unexpected nil response")
		return errors.New(constant.XrayRemoveOutBoundError)
	}
	return nil
}

// AddUser 添加用户
func AddUser(addUserDto dto.AddUserDto) error {
	hsClient := command.NewHandlerServiceClient(ClientConn)
	var resp *command.AlterInboundResponse
	switch addUserDto.Tag {
	case constant.ProtocolShadowsocks:
		resp, _ = hsClient.AlterInbound(context.Background(), &command.AlterInboundRequest{
			Tag: addUserDto.Tag,
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: addUserDto.Email,
						Account: serial.ToTypedMessage(&shadowsocks.Account{
							Password: addUserDto.SSPassword,
						}),
					},
				}),
		})
	case constant.ProtocolTrojan:
		resp, _ = hsClient.AlterInbound(context.Background(), &command.AlterInboundRequest{
			Tag: addUserDto.Tag,
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: addUserDto.Email,
						Account: serial.ToTypedMessage(&trojan.Account{
							Password: addUserDto.TrojanPassword,
						}),
					},
				}),
		})
	case constant.ProtocolVless:
		resp, _ = hsClient.AlterInbound(context.Background(), &command.AlterInboundRequest{
			Tag: addUserDto.Tag,
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: addUserDto.Email,
						Account: serial.ToTypedMessage(&vless.Account{
							Id: addUserDto.VId,
						}),
					},
				}),
		})
	case constant.ProtocolVmess:
		resp, _ = hsClient.AlterInbound(context.Background(), &command.AlterInboundRequest{
			Tag: addUserDto.Tag,
			Operation: serial.ToTypedMessage(
				&command.AddUserOperation{
					User: &protocol.User{
						Email: addUserDto.Email,
						Account: serial.ToTypedMessage(&vmess.Account{
							Id: addUserDto.VId,
						}),
					},
				}),
		})
	}
	if resp == nil {
		logrus.Errorf("nil response")
		return errors.New(constant.XrayAddUserError)
	}
	return nil
}

// RemoveUser 删除用户
func RemoveUser(tag string, email string) error {
	hsClient := command.NewHandlerServiceClient(ClientConn)
	resp, err := hsClient.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag:       tag,
		Operation: serial.ToTypedMessage(&command.RemoveUserOperation{Email: email}),
	})
	if err != nil {
		logrus.Errorf("删除用户异常 err: %v\n", err)
		return errors.New(constant.XrayRemoveUserError)
	}
	if resp == nil {
		logrus.Errorf("nil response")
		return errors.New(constant.XrayRemoveUserError)
	}
	return nil
}
