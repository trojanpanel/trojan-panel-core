package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"time"
	"trojan-panel-core/core"
)

var pool *redis.Pool

func InitRedis() {
	redisConfig := core.Config.RedisConfig
	pool = &redis.Pool{
		MaxIdle:     redisConfig.MaxIdle,
		MaxActive:   redisConfig.MaxActive,
		Wait:        redisConfig.Wait,
		IdleTimeout: 30 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
				redis.DialPassword(redisConfig.Password),
				redis.DialDatabase(redisConfig.Db),
			)
			if err != nil {
				logrus.Errorf("Redis初始化失败 err: %v", err)
				panic(err)
			}
			result, err := redis.String(conn.Do("PING"))
			if err != nil || result != "PONG" {
				conn.Close()
				logrus.Errorf("Redis连接失败 err: %v", err)
				panic(err)
			}
			return conn, nil
		},
	}
}

func CloseRedis() {
	if err := pool.Close(); err != nil {
		logrus.Errorf("redis close err: %v", err)
	}
}
