package cmd

import (
	"fmt"
	"os"

	"github.com/gregorychen3/rebalance-tool/internal/rebalance"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "rebalance",
	Long: "Interactive CLI utility to realign weightings of portfolio assets\nhttps://github.com/gregorychen3/rebalance-tool",
	Run: func(cmd *cobra.Command, args []string) {
		println("Enter desired weightings (in %):")
		dom := promptFloatInput("    Dom stock? ")
		intl := promptFloatInput("    Intl stock? ")
		bond := promptFloatInput("    Bond? ")
		weights := rebalance.NewTargetAlloc(dom, intl, bond)
		fmt.Printf("%v", weights)
	},
}

func promptFloatInput(msg string) float32 {
	fmt.Printf("%s", msg)

	var f float32
	if _, err := fmt.Scanf("%f", &f); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	return f
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
