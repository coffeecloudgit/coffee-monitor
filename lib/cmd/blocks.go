package cmd

import (
	"coffee-monitor/lib/log"
	"coffee-monitor/lib/shell"
	"fmt"

	"github.com/spf13/cobra"
)

var blocksCmd = &cobra.Command{
	Use:   "blocks",
	Short: "Print the mine blocks of coffee-cli",
	Long:  `Please Use command: coffee-cli blocks`,
	Run: func(cmd *cobra.Command, args []string) {
		err, rewardMap := shell.LotusMinerInfoGetRewardForLastBlocks()
		if err != nil {
			log.Logger.Error(err.Error())
			fmt.Println(err.Error())
		}
		log.Logger.Info("rewardMap:", "map:", rewardMap)
		fmt.Println(rewardMap)
	},
}

func init() {
	rootCmd.AddCommand(blocksCmd)
}
