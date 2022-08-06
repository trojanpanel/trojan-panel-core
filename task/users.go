package task

import (
	"github.com/sirupsen/logrus"
	"sync"
	"trojan-panel-core/dao"
)

func HandlerUsersDownloadAndUpload() {
	var mutex sync.Mutex
	defer mutex.Unlock()
	if mutex.TryLock() {

	}
}

func HandlerUsers() {
	var mutex sync.Mutex
	defer mutex.Unlock()
	if mutex.TryLock() {
		if err := dao.DeleteUsersByQuota(); err != nil {
			logrus.Errorf("删除流量超额用户失败 err: %v\n", err)
		}
	}
}
