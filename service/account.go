package service

import (
	"context"
	"github.com/avast/retry-go"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
	"trojan-core/dao"
	"trojan-core/model/constant"
	"trojan-core/proxy"
)

func HandleAccount() {
	go proxy.XrayCmdMap.Range(func(key, value any) bool {
		handleXrayAccountTraffic(key.(string))
		handleXrayAccountAuth(key.(string))
		return true
	})
	go proxy.HysteriaCmdMap.Range(func(key, value any) bool {
		handleHysteriaAccountTraffic(key.(string))
		return true
	})
	go proxy.NaiveProxyCmdMap.Range(func(key, value any) bool {
		handleNaiveProxyAccountTraffic(key.(string))
		handleNaiveProxyAccountAuth(key.(string))
		return true
	})
}

func XAddAccountTraffic(username string, tx int64, rx int64) {
	if err := retry.Do(func() error {
		_, err := dao.RedisClient.XAdd(context.Background(), &redis.XAddArgs{
			Stream: constant.AccountTrafficStream,
			ID:     "*",
			Values: map[string]interface{}{
				"username": username,
				"tx":       tx,
				"rx":       rx,
			},
		}).Result()
		return err
	}, []retry.Option{
		retry.Delay(3 * time.Second),
		retry.Attempts(2),
	}...); err != nil {
		logrus.Errorf("xadd account traffic err: %v", err)
	}
}
