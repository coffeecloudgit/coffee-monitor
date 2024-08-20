package cmd

import (
	"coffee-monitor/lib/log"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "coffee-cli",
	Short: "coffee-cli is a FIL monitor tool",
	Long:  `coffee-cli is a FIL monitor tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("coffee-cli root")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Logger.Info(err.Error())
		os.Exit(1)
	}
}
