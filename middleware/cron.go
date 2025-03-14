package middleware

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"time"
	"trojan-core/service"
)

func InitCron() error {
	loc := time.Now().Location()
	c := cron.New(cron.WithLocation(loc))
	_, err := c.AddFunc("@every 30s", service.HandleAccount)
	if err != nil {
		logrus.Errorf("cron add func HandleAccount err: %v", err)
		return fmt.Errorf("cron add func HandleAccount err")
	}
	c.Start()
	return nil
}
