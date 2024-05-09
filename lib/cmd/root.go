package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "coffee-cli",
	Short: "coffee-cli is a FIL monitor tool",
	Long:  `coffee-cli is a FIL monitor tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run coffee-cli...")
		//lib.AnalysisLog()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
