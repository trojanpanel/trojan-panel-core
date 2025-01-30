package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"trojan-panel-core/api"
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
	if err := api.StarServer(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
