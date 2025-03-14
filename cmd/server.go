package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"trojan-core/api"
	"trojan-core/dao"
	"trojan-core/middleware"
	"trojan-core/model/constant"
	"trojan-core/proxy"
	"trojan-core/router"
	"trojan-core/util"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Long:  "Run server.",
	Run:   runServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func runServer(cmd *cobra.Command, args []string) {
	defer releaseResource()

	middleware.InitLog()
	if err := initFile(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := proxy.InitProxy(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := middleware.InitCron(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	go func() {
		if err := api.StarGrpcServer(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	if err := startWebServer(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func startWebServer() error {
	r := gin.Default()
	router.Router(r)
	return r.Run(fmt.Sprintf(":%s", os.Getenv(constant.WebPort)))
}

func releaseResource() {
	if err := dao.CloseRedis(); err != nil {
		logrus.Errorf("release resource err: %v", err)
	}
}

func initFile() error {
	var dirs = []string{constant.LogDir, constant.BinDir,
		constant.XrayConfigDir, constant.SingBoxConfigDir,
		constant.HysteriaConfigDir, constant.NaiveProxyConfigDir,
	}
	for _, item := range dirs {
		if !util.Exists(item) {
			if err := os.Mkdir(item, os.ModePerm); err != nil {
				logrus.Errorf("%s create err: %v", item, err)
				return fmt.Errorf("%s create err", item)
			}
		}
	}
	return nil
}
