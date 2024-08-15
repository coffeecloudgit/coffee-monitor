package cmd

import (
	"coffee-monitor/lib"
	"coffee-monitor/lib/client"
	"coffee-monitor/lib/log"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "coffee-cli",
	Short: "coffee-cli is a FIL monitor tool",
	Long:  `coffee-cli is a FIL monitor tool`,
	Run: func(cmd *cobra.Command, args []string) {
		//run test wss server
		if client.IsConnectLocalhostServer() {
			log.Logger.Info("run test wss server")
			go client.RunTestServer()
			time.Sleep(5000 * time.Millisecond)
		}

		log.Logger.Info("run timer...")

		log.Logger.Info("connect server...")
		go client.ConnectServer()
		log.Logger.Info("connect server sleep 2s...")
		time.Sleep(2000 * time.Millisecond)
		//监控相关
		lib.Snapshot()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Logger.Info(err.Error())
		os.Exit(1)
	}
}
