package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of coffee-cli",
	Long:  `All software has versions. This is CoffeeMonitor's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("coffee-cli Generator v0.1 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
