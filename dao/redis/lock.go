package redis

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/sirupsen/logrus"
)

func RsLock(mutexName string) (*redsync.Mutex, error) {
	mutex := rs.NewMutex(mutexName)
	if err := mutex.Lock(); err != nil {
		logrus.Errorf("lock failed mutex name: %s err: %v", mutex.Name(), err)
		return nil, err
	}
	return mutex, nil
}

func RsUnLock(mutex *redsync.Mutex) {
	if ok, err := mutex.Unlock(); mutex == nil || !ok || err != nil {
		logrus.Errorf("unlock failed mutex name: %s", mutex.Name())
	}
}
