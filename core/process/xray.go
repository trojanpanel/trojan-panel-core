package process

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"regexp"
	"runtime"
	"sync"
	"trojan-panel-core/app/xray"
	"trojan-panel-core/module/constant"
	"trojan-panel-core/module/dto"
	"trojan-panel-core/service"
	"trojan-panel-core/util"
)

var userLinkRegex = regexp.MustCompile("user>>>([^>]+)>>>traffic>>>(downlink|uplink)")

type XrayProcess struct {
	process
}

func NewXrayProcess(apiPort uint) (*XrayProcess, error) {
	var mutex sync.Mutex
	defer mutex.Unlock()
	if mutex.TryLock() {
		x := &XrayProcess{process{binaryType: 1}}
		binaryFilePath, err := util.GetBinaryFile(1)
		if err != nil {
			return nil, err
		}
		configFilePath, err := util.GetConfigFile(1, apiPort)
		if err != nil {
			return nil, err
		}
		cmd := exec.Command(binaryFilePath, "-c", configFilePath)
		x.cmdMap.Store(0, cmd)
		runtime.SetFinalizer(x, x.Stop(apiPort))
		return x, nil
	}
	logrus.Errorf("new xray process errror err: lock not acquired\n")
	return nil, errors.New(constant.NewXrayProcessError)
}

func (x *XrayProcess) StartXray(apiPort uint) error {
	defer x.mutex.Unlock()
	if x.mutex.TryLock() {
		if x.IsRunning(apiPort) {
			return nil
		}
		cmd, ok := x.cmdMap.Load(apiPort)
		if ok {
			if err := cmd.(*exec.Cmd).Start(); err != nil {
				logrus.Errorf("start xray error err: %v\n", err)
				return errors.New(constant.XrayStartError)
			}
			go x.handlerUserUploadAndDownload(apiPort)
			go x.handlerUsers(apiPort)
			return nil
		}
		logrus.Errorf("start xray error err: process not found\n")
		return errors.New(constant.XrayStartError)
	}
	logrus.Errorf("start xray error err: lock not acquired\n")
	return errors.New(constant.XrayStartError)
}

func (x *XrayProcess) handlerUserUploadAndDownload(apiPort uint) {
	api := xray.NewXrayApi(apiPort)
	for {
		if !x.IsRunning(apiPort) {
			logrus.Errorf("数据库同步至Xray apiPort: %d xray not running\n", apiPort)
			break
		}
		addUserApiVos, err := service.SelectUsersToApi(true)
		if err != nil {
			logrus.Errorf("数据库同步至Xray apiPort: %d 查询用户失败 err: %v\n", apiPort, err)
		} else {
			for _, apiUserVo := range addUserApiVos {
				// 如果应用中存在则跳过
				stats, err := api.GetUserStats(apiUserVo.Password, "downlink", false)
				if err != nil || stats != nil {
					continue
				}
				userDto := dto.XrayAddUserDto{
					Tag:   "user",
					Email: apiUserVo.Password,
				}
				if err := api.AddUser(userDto); err != nil {
					logrus.Errorf("数据库同步至Xray apiPort: %d 添加用户失败 err: %v", apiPort, err)
					continue
				}
			}
		}
		apiUserVos, err := service.SelectUsersToApi(false)
		if err != nil {
			logrus.Errorf("数据库同步至Xray apiPort: %d 查询用户失败 err: %v\n", apiPort, err)
		} else {
			for _, apiUser := range apiUserVos {
				if err := api.DeleteUser("user", apiUser.Password); err != nil {
					logrus.Errorf("数据库同步至Xray apiPort: %d 删除用户失败 err: %v", apiPort, err)
					continue
				}
			}
		}
	}
}

func (x *XrayProcess) handlerUsers(apiPort uint) {
	api := xray.NewXrayApi(apiPort)
	for {
		if !x.IsRunning(apiPort) {
			logrus.Errorf("数据库同步至Xray apiPort: %d xray not running\n", apiPort)
			break
		}
		stats, err := api.QueryStats("", false)
		if err != nil {
			continue
		}
		users := make([]dto.UsersUpdateDto, 0)
		for _, stat := range stats {
			submatch := userLinkRegex.FindStringSubmatch(stat.Name)
			updateDto := dto.UsersUpdateDto{}
			if len(submatch) > 0 {
				email := submatch[0]
				isDown := submatch[1] == "downlink"
				updateDto.Password = email
				if isDown {
					updateDto.Download = stat.Value
				} else {
					updateDto.Upload = stat.Value
				}
				users = append(users, updateDto)
			}
		}
		for _, user := range users {
			encodePassword, err := util.AesEncode(user.Password)
			if err != nil {
				continue
			}
			download := user.Download
			upload := user.Upload
			if err := service.UpdateUser(nil, &apiPort, &encodePassword, &download,
				&upload); err != nil {
				logrus.Errorf("Xray同步至数据库 apiPort: %d 更新用户失败 err: %v", apiPort, err)
				continue
			}
		}
	}
}
