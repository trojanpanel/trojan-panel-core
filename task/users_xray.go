package task

import (
	"github.com/sirupsen/logrus"
	"sync"
	"trojan-panel-core/dao"
	"trojan-panel-core/module"
)

func HandlerUsersXrayDownloadAndUpload() {
	var mutex sync.Mutex
	tryLock := mutex.TryLock()
	if tryLock {

		usersXray := module.UsersXray{
			Password: nil,
			Download: nil,
			Upload:   nil,
		}
		if err := dao.UpdateUser(&usersXray); err != nil {
			logrus.Errorf("更新用户流量失败 err: %v\n", err)
		}
	}
	defer mutex.Unlock()
}

func HandlerUsersXrayStatus() {
	var mutex sync.Mutex
	if mutex.TryLock() {
		if err := dao.DeleteUsersByQuota(); err != nil {
			logrus.Errorf("删除流量不足用户失败 err: %v\n", err)
		}
	}
	defer mutex.Unlock()
}
