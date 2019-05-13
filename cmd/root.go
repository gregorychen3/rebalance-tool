package cmd

import (
	"fmt"
	"os"

	"github.com/gregorychen3/rebalance-tool/cmd/config"
	"github.com/gregorychen3/rebalance-tool/internal/portfolio"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "rebalance [config_file]",
	Long: "Interactive CLI utility to realign weightings of portfolio assets.\nhttps://github.com/gregorychen3/rebalance-tool",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		weightsMap, err := config.ReadConfigFile(args[0])
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		targetAllocation, err := portfolio.NewAssetAlloc(weightsMap)
		if err != nil {
			println(errors.Wrap(err, "failed creating target AssetAlloc").Error())
			os.Exit(1)
		}

		curHoldings := promptCurHoldings(targetAllocation)

		fmt.Printf("Adjustments to rebalance to unchanged portfolio value of $%.2f:\n", curHoldings.Total())
		rebalanceReport := portfolio.NewRebalanceReport(targetAllocation, curHoldings.Total(), curHoldings)
		prettyPrintReport(rebalanceReport)

		topupTotal := portfolio.TopupTotal(targetAllocation, curHoldings)
		fmt.Printf("Adjustments to \"top-up\" rebalance to new portfolio value of $%.2f:\n", topupTotal)
		topupRebalanceReport := portfolio.NewRebalanceReport(targetAllocation, topupTotal, curHoldings)
		prettyPrintReport(topupRebalanceReport)
	},
}

func promptCurHoldings(alloc portfolio.AssetAlloc) portfolio.Holdings {
	println("Enter current portfolio holdings (in $):")
	holdings := portfolio.Holdings{}
	for k := range alloc {
		holdings[k] = promptFloatInput(fmt.Sprintf("    %v? ", k))
	}
	return holdings
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

func prettyPrintReport(r portfolio.RebalanceReport) {
	for k, v := range r {
		fmt.Printf("    %v:  %+.2f\n", k, v)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
