package cmd

import (
	"coffee-monitor/lib"
	"coffee-monitor/lib/client"
	"coffee-monitor/lib/log"
	"github.com/spf13/cobra"
	"time"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run coffee monitor",
	Long:  `Run coffee monitor. example: coffee-monitor run`,
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

		select {} // 阻塞主 goroutine，防止程序退出
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
