package naiveproxy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"trojan-panel-core/module/bo"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
)

type naiveProxyApi struct {
	apiPort uint
}

// NewNaiveProxyApi 初始化NaiveProxy Api
func NewNaiveProxyApi(apiPort uint) *naiveProxyApi {
	return &naiveProxyApi{
		apiPort: apiPort,
	}
}

// ListUsers 查询节点上的所有用户
func (n *naiveProxyApi) ListUsers() (*[]bo.HandleAuth, error) {
	url := fmt.Sprintf("http://127.0.0.1:%d/config/apps/http/servers/srv0/routes/0/handle/0/routes/0/handle/", n.apiPort)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Errorf("NaiveProxy ListUsers NewRequest err: %v", err)
		return nil, errors.New(constant.SysError)
	}
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode != 200 {
		logrus.Errorf("NaiveProxy ListUsers http resp err: %v", err)
		return nil, errors.New(constant.SysError)
	}
	contentByte, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("NaiveProxy ListUsers IO err: %v", err)
		return nil, errors.New(constant.SysError)
	}
	var handleAuths *[]bo.HandleAuth
	if err = json.Unmarshal(contentByte, &handleAuths); err != nil {
		logrus.Errorf("NaiveProxy ListUsers Unmarshal err: %v", err)
		return nil, errors.New(constant.SysError)
	}
	return handleAuths, nil
}

// GetUser 查询节点上的用户
func (n *naiveProxyApi) GetUser(pass string) (*bo.HandleAuth, *int, error) {
	users, err := n.ListUsers()
	if err != nil {
		return nil, nil, err
	}
	for index, user := range *users {
		if user.AuthPassDeprecated == pass {
			return &user, &index, nil
		}
	}
	return nil, nil, nil
}

// AddUser 节点上添加用户
func (n *naiveProxyApi) AddUser(dto dto.NaiveProxyAddUserDto) error {
	user, _, err := n.GetUser(dto.Pass)
	if err != nil {
		return err
	}
	if user != nil {
		return nil
	}

	authJsonStr := `{
    "handler":"forward_proxy",
    "hide_ip":true,
    "hide_via":true,
    "probe_resistance":{}
}`
	var handleAuth *bo.HandleAuth
	if err = json.Unmarshal([]byte(authJsonStr), &handleAuth); err != nil {
		logrus.Errorf("NaiveProxy AddUser Unmarshal err: %v", err)
		return errors.New(constant.SysError)
	}
	handleAuth.AuthUserDeprecated = dto.Username
	handleAuth.AuthPassDeprecated = dto.Pass
	addUserDtoByte, err := json.Marshal(handleAuth)
	if err != nil {
		logrus.Errorf("NaiveProxy AddUser Marshal err: %v", err)
		return errors.New(constant.SysError)
	}

	url := fmt.Sprintf("http://127.0.0.1:%d/config/apps/http/servers/srv0/routes/0/handle/0/routes/0/handle/0", n.apiPort)
	req, err := http.NewRequest("POST", url,
		bytes.NewBuffer(addUserDtoByte))
	if err != nil {
		logrus.Errorf("NaiveProxy AddUser NewRequest err: %v", err)
		return errors.New(constant.SysError)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode != 200 {
		logrus.Errorf("NaiveProxy AddUser resp err: %v", err)
		return errors.New(constant.SysError)
	}
	return nil
}

// DeleteUser 节点上删除用户
func (n *naiveProxyApi) DeleteUser(pass string) error {
	_, index, err := n.GetUser(pass)
	if err != nil {
		return err
	}
	if index != nil {
		url := fmt.Sprintf("http://127.0.0.1:%d/config/apps/http/servers/srv0/routes/0/handle/0/routes/0/handle/%d", n.apiPort, *index)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			logrus.Errorf("NaiveProxy DeleteUser NewRequest err: %v", err)
			return errors.New(constant.SysError)
		}
		resp, err := http.DefaultClient.Do(req)
		defer func() {
			if resp != nil {
				resp.Body.Close()
			}
		}()
		if err != nil || resp.StatusCode != 200 {
			logrus.Errorf("NaiveProxy DeleteUser resp err: %v", err)
			return errors.New(constant.SysError)
		}
		return nil
	}
	return nil
}
