package middleware

import (
	"github.com/robfig/cron/v3"
	"time"
	"xray-manage/job"
)

// InitCron 初始化定时任务
func InitCron() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithLocation(location))
	c.AddFunc("@every 5s", job.HandlerUsersXrayDownloadAndUpload)
	c.AddFunc("@every 10s", job.HandlerUsersXrayStatus)
	c.Start()
}
