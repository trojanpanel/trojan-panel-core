package xray

import (
	"fmt"
	"regexp"
	"testing"
	"trojan-panel-core/model/dto"
)

var userLinkRegex = regexp.MustCompile("user>>>([^>]+)>>>traffic>>>(downlink|uplink)")

func TestXrayListUsers(t *testing.T) {
	api := NewXrayApi(30451)
	xrayStatsVos, err := api.QueryStats("", false)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for _, stat := range xrayStatsVos {
		submatch := userLinkRegex.FindStringSubmatch(stat.Name)
		if len(submatch) == 3 {
			fmt.Println(stat.Name)
			fmt.Println(stat.Value)
		}
	}
}

func TestXrayAddUser(t *testing.T) {
	api := NewXrayApi(30451)
	api.AddUser(dto.XrayAddUserDto{
		Protocol: "vless",
		Password: "123123",
	})
	TestXrayListUsers(t)
}
