package proxy

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
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
func (n *NaiveProxyApi) ListUsers() ([]string, error) {
	var authCredentials []string
	if !NewHysteriaInstance(n.apiPort).IsRunning() {
		return authCredentials, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := fmt.Sprintf("http://127.0.0.1:%s/config/apps/http/servers/srv0/routes/0/handle/0/routes/0/handle/0/auth_credentials/", n.apiPort)
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
	if err = json.Unmarshal(body, &authCredentials); err != nil {
		logrus.Errorf("NaiveProxy ListUsers Unmarshal err: %v", err)
		return nil, fmt.Errorf(constant.SysError)
	}
	return authCredentials, nil
}

// HandleUser add or delete user on node
func (n *NaiveProxyApi) HandleUser(authCredentials []string, add bool) error {
	if !NewHysteriaInstance(n.apiPort).IsRunning() {
		return nil
	}
	authCredentialsByte, err := json.Marshal(authCredentials)
	if err != nil {
		logrus.Errorf("NaiveProxy HandleUser Marshal err: %v", err)
		return fmt.Errorf(constant.SysError)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := fmt.Sprintf("http://127.0.0.1:%s/config/apps/http/servers/srv0/routes/0/handle/0/routes/0/handle/0/auth_credentials/", n.apiPort)
	method := http.MethodPost
	if !add {
		method = http.MethodDelete
	}
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(authCredentialsByte))
	if err != nil {
		logrus.Errorf("NaiveProxy HandleUser NewRequest err: %v", err)
		return fmt.Errorf(constant.SysError)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode != http.StatusOK {
		logrus.Errorf("NaiveProxy HandleUser http resp err: %v", err)
		return fmt.Errorf(constant.HttpError)
	}
	return nil
}

func handleNaiveProxyAuthCredential(user, pass string) string {
	return base64.StdEncoding.EncodeToString([]byte(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, pass)))))
}
