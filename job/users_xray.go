package job

import (
	"github.com/sirupsen/logrus"
	"sync"
	"xray-manage/dao"
	"xray-manage/module"
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
		if err := dao.UpdateUsersXray(&usersXray); err != nil {
			logrus.Errorf("更新用户流量失败 err: %v\n", err)
		}
	}
	defer mutex.Unlock()
}

func HandlerUsersXrayStatus() {
	var mutex sync.Mutex
	if mutex.TryLock() {
		if err := dao.DeleteUsersXrayByQuota(); err != nil {
			logrus.Errorf("删除流量不足用户失败 err: %v\n", err)
		}
	}
	defer mutex.Unlock()
}
