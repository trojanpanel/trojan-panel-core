package proxy

import (
	"errors"
)

func InitProxy() error {
	return nil
}

func StartProxy(proxy string, key string, value []byte) error {
	return errors.New("proxy not supported")
}

func StopProxy(proxy string, key string) error {
	return errors.New("proxy not supported")
}
