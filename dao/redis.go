package dao

import "github.com/redis/go-redis/v9"

var RedisClient *redis.Client

func InitRedis(addr string, password string) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})
}

func CloseRedis() error {
	return RedisClient.Close()
}
