package main

import (
	"fmt"
	"github.com/xtls/xray-core/app/commander"
	"github.com/xtls/xray-core/app/policy"
	"github.com/xtls/xray-core/app/router"
	"github.com/xtls/xray-core/app/stats"
	statscmd "github.com/xtls/xray-core/app/stats/command"
	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/proxy/dokodemo"
	"github.com/xtls/xray-core/proxy/freedom"
	"testing"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/pkg/trojango"
	"trojan-panel-core/pkg/xray"
)

func TestStartXray(t *testing.T) {
	config := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&stats.Config{}),
			serial.ToTypedMessage(&commander.Config{
				Tag: "api",
				Service: []*serial.TypedMessage{
					serial.ToTypedMessage(&statscmd.Config{}),
				},
			}),
			serial.ToTypedMessage(&policy.Config{
				Level: map[uint32]*policy.Policy{
					0: {
						Stats: &policy.Policy_Stats{
							UserUplink:   true,
							UserDownlink: true,
						},
					},
				},
				System: &policy.SystemPolicy{
					Stats: &policy.SystemPolicy_Stats{
						InboundUplink:    true,
						InboundDownlink:  true,
						OutboundUplink:   true,
						OutboundDownlink: true,
					},
				},
			}),
			serial.ToTypedMessage(&router.Config{
				Rule: []*router.RoutingRule{
					{
						InboundTag: []string{"api"},
						TargetTag: &router.RoutingRule_Tag{
							Tag: "api",
						},
					},
				},
			}),
		},
		Inbound: []*core.InboundHandlerConfig{
			{
				Tag: "api",
				ProxySettings: serial.ToTypedMessage(&dokodemo.Config{
					Address:  net.NewIPOrDomain(net.LocalHostIP),
					Port:     uint32(60087),
					Networks: []net.Network{net.Network_TCP},
				}),
			},
		},
		Outbound: []*core.OutboundHandlerConfig{
			{
				Tag:           "direct",
				ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
			},
		},
	}
	err := xray.StartXray(config)
	if err != nil {
		fmt.Println(err)
	}
}

func TestQueryStats(t *testing.T) {
	api := xray.XrayApi()
	userDto := dto.AddUserDto{
		Tag:            "trojan",
		TrojanPassword: "123123",
	}

	err := api.AddUser(userDto)
	if err != nil {
		fmt.Println(err)
	}

	stats, err := api.QueryStats("", false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stats)
}

func TestStartTrojanGo(t *testing.T) {
	trojango.StartTrojanGo()
}
