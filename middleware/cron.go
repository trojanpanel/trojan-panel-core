package middleware

import (
	"errors"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"time"
	"trojan-core/app"
)

func InitCron() error {
	location, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithLocation(location))
	_, err := c.AddFunc("@every 50s", app.CronHandlerUser)
	if err != nil {
		logrus.Errorf("cron add func CronHandlerUser err: %v", err)
		return errors.New("cron add func CronHandlerUser err")
	}
	_, err = c.AddFunc("@every 50s", app.CronHandlerDownloadAndUpload)
	if err != nil {
		logrus.Errorf("cron add func CronHandlerDownloadAndUpload err: %v", err)
		return errors.New("cron add func CronHandlerDownloadAndUpload err")
	}
	c.Start()
	return nil
}
