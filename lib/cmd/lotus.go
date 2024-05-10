package cmd

import (
	"coffee-monitor/lib/fil"
	"github.com/spf13/cobra"
	"log"
)

var lotusCmd = &cobra.Command{
	Use:   "lotus",
	Short: "Show Lotus Node Info",
	Long:  `Show Lotus Node Info. example: coffee-monitor lotus`,
	Run: func(cmd *cobra.Command, args []string) {
		lotusInfo, err := fil.GetLotusInfo()
		if err != nil {
			log.Fatal(err)
		}

		log.Println(lotusInfo)
	},
}

func init() {
	rootCmd.AddCommand(lotusCmd)
}
