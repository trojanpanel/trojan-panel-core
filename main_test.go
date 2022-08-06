package main

import (
	"fmt"
	"testing"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/pkg/xray"
)

func TestStartXray(t *testing.T) {
	xray.StartXray()
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
