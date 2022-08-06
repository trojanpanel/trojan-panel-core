package middleware

import (
	"github.com/robfig/cron/v3"
	"time"
	"trojan-panel-core/task"
)

// InitCron 初始化定时任务
func InitCron() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithLocation(location))
	c.AddFunc("@every 10s", task.HandlerUsersDownloadAndUpload)
	c.AddFunc("@every 15s", task.HandlerUsers)
	c.AddFunc("@every 20s", task.HandlerUsersXray)
	c.AddFunc("@every 20s", task.HandlerUsersTrojanGo)
	c.AddFunc("@every 20s", task.HandlerUsersTrojanGo)
	c.Start()
}
