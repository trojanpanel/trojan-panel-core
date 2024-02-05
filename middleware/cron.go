package middleware

import (
	"github.com/robfig/cron/v3"
	"time"
	"trojan-panel-core/app"
)

// InitCron initialize scheduled tasks
func InitCron() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithLocation(location))
	_, _ = c.AddFunc("@every 50s", app.CronHandlerUser)
	_, _ = c.AddFunc("@every 50s", app.CronHandlerDownloadAndUpload)
	c.Start()
}
