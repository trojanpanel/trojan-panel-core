package proxy

import (
	"bytes"
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

type HysteriaApi struct {
	apiPort string
}

func NewHysteriaApi(apiPort string) *HysteriaApi {
	return &HysteriaApi{
		apiPort: apiPort,
	}
}

// ListUsers 每个用户的流量信息
func (h *HysteriaApi) ListUsers(clear bool, secret string) (map[string]bo.HysteriaUserTraffic, error) {
	var users map[string]bo.HysteriaUserTraffic
	if !NewHysteriaInstance(h.apiPort).IsRunning() {
		return users, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := fmt.Sprintf("http://127.0.0.1:%s/traffic", h.apiPort)
	if clear {
		url = fmt.Sprintf("%s?clear=1", url)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logrus.Errorf("Hysteria ListUsers NewRequest err: %v", err)
		return nil, fmt.Errorf(constant.SysError)
	}
	req.Header.Set("Authorization", secret)
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode != http.StatusOK {
		logrus.Errorf("Hysteria ListUsers err: %v", err)
		return nil, fmt.Errorf(constant.HttpError)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Hysteria io read err: %v", err)
		return nil, fmt.Errorf(constant.HttpError)
	}
	if err = json.Unmarshal(body, &users); err != nil {
		logrus.Errorf("Hysteria ListUsers Unmarshal err: %v", err)
		return nil, fmt.Errorf(constant.SysError)
	}
	return users, nil
}

// KickUsers 踢下线
func (h *HysteriaApi) KickUsers(keys []string, secret string) error {
	if !NewHysteriaInstance(h.apiPort).IsRunning() {
		return nil
	}
	usernamesByte, err := json.Marshal(keys)
	if err != nil {
		logrus.Errorf("Hysteria KickUsers Marshal err: %v", err)
		return fmt.Errorf(constant.SysError)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := fmt.Sprintf("http://127.0.0.1:%s/kick", h.apiPort)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url,
		bytes.NewBuffer(usernamesByte))
	if err != nil {
		logrus.Errorf("Hysteria KickUsers NewRequest err: %v", err)
		return fmt.Errorf(constant.SysError)
	}
	req.Header.Set("Authorization", secret)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode != http.StatusOK {
		logrus.Errorf("Hysteria KickUsers err: %v", err)
		return fmt.Errorf(constant.HttpError)
	}
	return nil
}

// OnlineUsers 在线用户
func (h *HysteriaApi) OnlineUsers(secret string) (map[string]int64, error) {
	var onlineUsers map[string]int64
	if !NewHysteriaInstance(h.apiPort).IsRunning() {
		return onlineUsers, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := fmt.Sprintf("http://127.0.0.1:%s/online", h.apiPort)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		logrus.Errorf("Hysteria OnlineUsers NewRequest err: %v", err)
		return nil, fmt.Errorf(constant.SysError)
	}
	req.Header.Set("Authorization", secret)
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode != http.StatusOK {
		logrus.Errorf("Hysteria OnlineUsers err: %v", err)
		return nil, fmt.Errorf(constant.HttpError)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Hysteria io read err: %v", err)
		return nil, fmt.Errorf(constant.HttpError)
	}
	if err = json.Unmarshal(body, &onlineUsers); err != nil {
		logrus.Errorf("Hysteria OnlineUsers Unmarshal err: %v", err)
		return nil, fmt.Errorf(constant.SysError)
	}
	return onlineUsers, nil
}
