package cmd

import (
	"coffee-monitor/lib"
	"coffee-monitor/lib/fil"
	"coffee-monitor/lib/log"
	"fmt"
	"github.com/spf13/cobra"
)

var lotusCmd = &cobra.Command{
	Use:   "lotus",
	Short: "Show Lotus Node Info",
	Long:  `Show Lotus Node Info. example: coffee-monitor lotus`,
	Run: func(cmd *cobra.Command, args []string) {
		lotusInfo, err := fil.GetLotusInfo()
		if err != nil {
			log.Logger.Error(err.Error())
		}

		fmt.Println(lotusInfo)

		lib.Lotus()
	},
}

func init() {
	rootCmd.AddCommand(lotusCmd)
}
