package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
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
func (h *HysteriaApi) ListUsers(clear bool) (map[string]bo.HysteriaUserTraffic, error) {
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
	req.Header.Set("Authorization", os.Getenv(constant.HysteriaAuthSecret))
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil || resp.StatusCode != http.StatusOK {
		logrus.Errorf("Hysteria ListUsers http resp err: %v", err)
		return nil, fmt.Errorf(constant.HttpError)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Hysteria ListUsers io read err: %v", err)
		return nil, fmt.Errorf(constant.SysError)
	}
	if err = json.Unmarshal(body, &users); err != nil {
		logrus.Errorf("Hysteria ListUsers Unmarshal err: %v", err)
		return nil, fmt.Errorf(constant.SysError)
	}
	return users, nil
}
