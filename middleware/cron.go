package middleware

import (
	"github.com/robfig/cron/v3"
	"time"
)

func InitCron() error {
	loc := time.Now().Location()
	c := cron.New(cron.WithLocation(loc))
	c.Start()
	return nil
}
