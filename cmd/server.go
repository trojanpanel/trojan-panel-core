package cmd

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"trojan-core/dao"
	"trojan-core/middleware"
	"trojan-core/model/constant"
	"trojan-core/router"
	"trojan-core/util"
)

func runServer(serverPort string) error {
	defer releaseResource()

	middleware.InitLog()
	if err := initFile(); err != nil {
		return err
	}
	if err := dao.InitSqliteDB(); err != nil {
		return err
	}
	if err := middleware.InitCron(); err != nil {
		return err
	}

	r := gin.Default()
	router.Router(r)
	return r.Run(fmt.Sprintf(":%s", serverPort))
}

func releaseResource() {
	if err := dao.CloseSqliteDB(); err != nil {
		logrus.Errorf(err.Error())
	}
}

func initFile() error {
	var dirs = []string{constant.LogPath,
		constant.XrayPath, constant.HysteriaPath, constant.NaivePath,
		constant.XrayBinPath, constant.HysteriaBinPath, constant.NaiveBinPath,
	}
	for _, item := range dirs {
		if !util.Exists(item) {
			if err := os.Mkdir(item, os.ModePerm); err != nil {
				logrus.Errorf("%s create err: %v", item, err)
				return errors.New(fmt.Sprintf("%s create err", item))
			}
		}
	}
	return nil
}
