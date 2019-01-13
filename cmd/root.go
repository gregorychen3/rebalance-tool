package cmd

import (
	"fmt"
	"os"

	"github.com/gregorychen3/rebalance-tool/cmd/config"
	"github.com/gregorychen3/rebalance-tool/internal/portfolio"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const configFileFlag = "config"

var cfgFile string

func init() {
	rootCmd.Flags().StringVar(&cfgFile, configFileFlag, "rebalance_conf.json", "Config file defining target asset allocation")
	rootCmd.MarkFlagRequired(configFileFlag)
}

var rootCmd = &cobra.Command{
	Use:  "rebalance",
	Long: "Interactive CLI utility to realign weightings of portfolio assets.\nhttps://github.com/gregorychen3/rebalance-tool",
	Run: func(cmd *cobra.Command, args []string) {
		weightsMap, err := config.ReadConfigFile(cfgFile)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		targetAllocation, err := portfolio.NewAssetAllocation(weightsMap)
		if err != nil {
			println(errors.Wrap(err, "failed creating target AssetAllocation").Error())
			os.Exit(1)
		}

		curHoldings := promptCurHoldings(targetAllocation)

		fmt.Printf("%v\n", targetAllocation)
		fmt.Printf("%v\n", curHoldings)

		fmt.Printf("Adjustments to rebalance to unchanged portfolio value of $%.2f:\n", curHoldings.Total())
		/*
			rebalanceReport := portfolio.NewRebalanceReport(targetAlloc, curHoldings.Total(), curHoldings)
			prettyPrintReport(rebalanceReport)

			topupTotal := portfolio.TopupTotal(targetAlloc, curHoldings)
			fmt.Printf("Adjustments to \"top-up\" rebalance to new portfolio value of $%.2f:\n", topupTotal)
			topupRebalanceReport := portfolio.NewRebalanceReport(targetAlloc, topupTotal, curHoldings)
			prettyPrintReport(topupRebalanceReport)
		*/
	},
}

func promptTargetAlloc() *portfolio.AssetAlloc {
	println("Enter target asset allocations (in %):")
	dom := promptFloatInput("    Dom stock? ") / 100
	intl := promptFloatInput("    Intl stock? ") / 100
	bond := promptFloatInput("    Bond? ") / 100
	return portfolio.NewAssetAlloc(dom, intl, bond)
}

func promptCurHoldings(alloc portfolio.AssetAllocation) portfolio.Holdingss {

	println("Enter current portfolio holdings (in $):")
	holdings := portfolio.Holdingss{}
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

func prettyPrintReport(r *portfolio.RebalanceReport) {
	fmt.Printf("    Dom:  %+.2f\n    Intl: %+.2f\n    Bond: %+.2f\n", r.Dom, r.Intl, r.Bond)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
