package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
	"trojan-core/model/bo"
	"trojan-core/model/constant"
)

type NaiveProxyApi struct {
	apiPort string
}

func NewNaiveProxyApi(apiPort string) *NaiveProxyApi {
	return &NaiveProxyApi{
		apiPort: apiPort,
	}
}

// ListUsers query all users on a node
func (n *NaiveProxyApi) ListUsers() ([]bo.HandleAuth, error) {
	var handleAuths []bo.HandleAuth
	if !NewHysteriaInstance(n.apiPort).IsRunning() {
		return handleAuths, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := fmt.Sprintf("http://127.0.0.1:%s/config/apps/http/servers/srv0/routes/0/handle/0/routes/0/handle/", n.apiPort)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logrus.Errorf("NaiveProxy ListUsers NewRequest err: %v", err)
		return nil, fmt.Errorf(constant.SysError)
	}
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode != http.StatusOK {
		logrus.Errorf("NaiveProxy ListUsers http resp err: %v", err)
		return nil, fmt.Errorf(constant.HttpError)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("NaiveProxy ListUsers io read err: %v", err)
		return nil, fmt.Errorf(constant.SysError)
	}
	if err = json.Unmarshal(body, &handleAuths); err != nil {
		logrus.Errorf("NaiveProxy ListUsers Unmarshal err: %v", err)
		return nil, fmt.Errorf(constant.SysError)
	}
	return handleAuths, nil
}
