package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"trojan-core/util"
)

var port string

func init() {
	rootCmd.Flags().StringVarP(&port, "port", "p", "", "The port of the web server")
}

var rootCmd = &cobra.Command{
	Use:   "trojan-core",
	Short: "A command line tool for trojan-core",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	if err := util.VerifyPort(port); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for {
		if err := runServer(port); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
