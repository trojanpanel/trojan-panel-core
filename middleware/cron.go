package middleware

import (
	"github.com/robfig/cron/v3"
	"time"
	"trojan-panel-core/service"
)

// InitCron 初始化定时任务
func InitCron() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithLocation(location))
	c.AddFunc("@every 8s", service.CronHandlerUser)
	c.AddFunc("@every 10s", service.CronHandlerDownloadAndUpload)
	c.Start()
}
