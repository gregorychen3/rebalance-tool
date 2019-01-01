package cmd

import (
	"fmt"
	"os"

	"github.com/gregorychen3/rebalance-tool/internal/portfolio"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "rebalance",
	Long: "Interactive CLI utility to realign weightings of portfolio assets.\nhttps://github.com/gregorychen3/rebalance-tool",
	Run: func(cmd *cobra.Command, args []string) {
		targetAlloc := promptTargetAlloc()
		curHoldings := promptCurHoldings()

		fmt.Printf("Adjustments to rebalance to unchanged portfolio value of $%.2f:\n", curHoldings.Total())
		rebalanceReport := portfolio.NewRebalanceReport(targetAlloc, curHoldings.Total(), curHoldings)
		prettyPrintReport(rebalanceReport)

		topupTotal := portfolio.TopupTotal(targetAlloc, curHoldings)
		fmt.Printf("Adjustments to \"top-up\" rebalance to new portfolio value of $%.2f:\n", topupTotal)
		topupRebalanceReport := portfolio.NewRebalanceReport(targetAlloc, topupTotal, curHoldings)
		prettyPrintReport(topupRebalanceReport)
	},
}

func promptTargetAlloc() *portfolio.AssetAlloc {
	println("Enter target asset allocations (in %):")
	dom := promptFloatInput("    Dom stock? ") / 100
	intl := promptFloatInput("    Intl stock? ") / 100
	bond := promptFloatInput("    Bond? ") / 100
	return portfolio.NewAssetAlloc(dom, intl, bond)
}

func promptCurHoldings() *portfolio.Holdings {
	println("Enter current portfolio holdings (in $):")
	dom := promptFloatInput("    Dom stock? ")
	intl := promptFloatInput("    Intl stock? ")
	bond := promptFloatInput("    Bond? ")
	return portfolio.NewHoldings(dom, intl, bond)
}

func promptFloatInput(msg string) float64 {
	fmt.Printf("%s", msg)
	var f float64
	if _, err := fmt.Scanf("%f", &f); err != nil {
		println(err.Error())
		os.Exit(1)
	}
	return f
}

func prettyPrintReport(r *portfolio.RebalanceReport) {
	fmt.Printf("    Dom:  %+.2f\n    Intl: %+.2f\n    Bond: %+.2f\n", r.Dom, r.Intl, r.Bond)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
