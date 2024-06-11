package cmd

import (
	"coffee-monitor/lib"
	"coffee-monitor/lib/client"
	fil "coffee-monitor/lib/fil/miner"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "coffee-cli",
	Short: "coffee-cli is a FIL monitor tool",
	Long:  `coffee-cli is a FIL monitor tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run timer...")
		log.Println("connect server...")
		go client.ConnectServer()
		log.Println("connect server sleep 2s...")
		time.Sleep(2000 * time.Millisecond)
		lib.Snapshot()
		err := fil.MinerLogTailProcessor()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
