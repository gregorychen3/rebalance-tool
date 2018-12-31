package portfolio

import (
	"fmt"
)

type RebalanceReport struct {
	Dom  float64
	Intl float64
	Bond float64
}

func NewRebalanceReport(targetAlloc *AssetAlloc, targetTotal float64, curHoldings *Holdings) *RebalanceReport {
	targetHoldings := Holdings{
		dom:  targetTotal * targetAlloc.dom,
		intl: targetTotal * targetAlloc.intl,
		bond: targetTotal * targetAlloc.bond,
	}
	return &RebalanceReport{
		Dom:  targetHoldings.dom - curHoldings.dom,
		Intl: targetHoldings.intl - curHoldings.intl,
		Bond: targetHoldings.bond - curHoldings.bond,
	}
}

func (r *RebalanceReport) Pretty() string {
	return fmt.Sprintf("    Dom:  %+.2f\n    Intl: %+.2f\n    Bond: %+.2f\n", r.Dom, r.Intl, r.Bond)
}

func TopupTotal(targetAlloc *AssetAlloc, curHoldings *Holdings) float64 {
	// find most overweighted category
	curTotal := curHoldings.Total()
	curAlloc := NewAssetAlloc(curHoldings.dom/curTotal, curHoldings.intl/curTotal, curHoldings.bond/curTotal)

	allocDiffs := []float64{
		curAlloc.dom - targetAlloc.dom,
		curAlloc.intl - targetAlloc.intl,
		curAlloc.bond - targetAlloc.bond,
	}

	var topupTotal float64
	switch i := maxElementIndex(allocDiffs); i {
	case 0:
		topupTotal = curAlloc.dom / targetAlloc.dom
	case 1:
		topupTotal = curAlloc.intl / targetAlloc.intl
	case 2:
		topupTotal = curHoldings.bond / targetAlloc.bond
	}

	return topupTotal
}

func maxElementIndex(arr []float64) int {
	maxSeen := -1.0
	maxSeenIndex := -1
	for i, v := range arr {
		if v > maxSeen {
			maxSeen = v
			maxSeenIndex = i
		}
	}
	return maxSeenIndex
}
