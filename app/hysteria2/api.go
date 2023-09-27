package hysteria2

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
	"trojan-panel-core/model/bo"
	"trojan-panel-core/model/constant"
)

type hysteria2Api struct {
	apiPort uint
}

func NewHysteria2Api(apiPort uint) *hysteria2Api {
	return &hysteria2Api{
		apiPort: apiPort,
	}
}

func apiClient() *http.Client {
	return &http.Client{
		Timeout: 3 * time.Second,
	}
}

func (n *hysteria2Api) ListUsers(clear bool) (map[string]bo.Hysteria2UserTraffic, error) {
	client := apiClient()
	url := fmt.Sprintf("http://127.0.0.1:%d/traffic", n.apiPort)
	if clear {
		url = fmt.Sprintf("%s?clear=1", url)
	}
	resp, err := client.Get(url)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode != http.StatusOK {
		logrus.Errorf("Hysteria2 ListUsers err: %v", err)
		return nil, errors.New(constant.HttpError)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Hysteria2 io read err: %v", err)
		return nil, errors.New(constant.HttpError)
	}
	var users map[string]bo.Hysteria2UserTraffic
	if err = json.Unmarshal(body, &users); err != nil {
		logrus.Errorf("Hysteria2 ListUsers Unmarshal err: %v", err)
		return nil, errors.New(constant.SysError)
	}
	return users, nil
}

func (n *hysteria2Api) GetUser(pass string, clear bool) (*bo.Hysteria2User, error) {
	users, err := n.ListUsers(clear)
	if err != nil {
		return nil, err
	}
	user := users[pass]
	return &bo.Hysteria2User{
		Pass: pass,
		Tx:   user.Tx,
		Rx:   user.Rx,
	}, nil
}
