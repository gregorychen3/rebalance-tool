package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "rebalance",
	Long: "Interactive CLI utility to realign weightings of portfolio assets\nhttps://github.com/gregorychen3/rebalance-tool",
	Run: func(cmd *cobra.Command, args []string) {
		println("Run")
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
