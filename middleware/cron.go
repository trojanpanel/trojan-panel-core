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
	// 持续更新download upload字段
	c.AddFunc("@every 8s", task.HandlerUsersXrayDownloadAndUpload)
	// 删除 quota < download + upload
	c.AddFunc("@every 10s", task.HandlerUsersXrayStatus)
	c.Start()
}
