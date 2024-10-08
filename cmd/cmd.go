package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"trojan-core/model/constant"
)

var (
	config  string
	version bool
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "", "config file path")
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "show version")
}

var rootCmd = &cobra.Command{
	Use:   "trojan-core",
	Short: "the core of trojan panel",
	Long:  "the core of trojan panel",
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println("trojan-core version", constant.Version)
			os.Exit(0)
		}
		if config != "" {
			if err := runServer(); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
		fmt.Println("Usage: trojan-core [-c] [-v] [-h]")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
