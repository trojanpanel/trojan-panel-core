package dao

import (
	"github.com/redis/go-redis/v9"
	"os"
	"trojan-core/model/constant"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv(constant.RedisAddr),
		Password: os.Getenv(constant.RedisPass),
	})
}

func CloseRedis() error {
	return RedisClient.Close()
}
