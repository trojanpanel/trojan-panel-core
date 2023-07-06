package xray

import (
	"fmt"
	"regexp"
	"testing"
)

var userLinkRegex = regexp.MustCompile("user>>>([^>]+)>>>traffic>>>(downlink|uplink)")

func TestXrayListUsers(t *testing.T) {
	api := NewXrayApi(30452)
	xrayStatsVos, err := api.QueryStats("", false)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	for _, stat := range xrayStatsVos {
		submatch := userLinkRegex.FindStringSubmatch(stat.Name)
		if len(submatch) == 2 {
			fmt.Println(stat.Name)
			fmt.Println(stat.Value)
		}

	}
}
