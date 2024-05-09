package cmd

import (
	"coffee-monitor/lib"
	"github.com/spf13/cobra"
)

var checkLog string

var checkLogCmd = &cobra.Command{
	Use:   "check",
	Short: "Analysis the miner log",
	Long:  `Analysis the miner log,find forked block. example: coffee-monitor check`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.AnalysisLog(checkLog)
	},
}

func init() {
	checkLogCmd.Flags().StringVarP(&checkLog, "log", "l", "", "Add the miner log path,example: coffee-monitor check -l '/var/log/miner.log'")
	rootCmd.AddCommand(checkLogCmd)
}
