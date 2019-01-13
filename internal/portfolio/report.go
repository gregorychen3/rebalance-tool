package portfolio

type RebalanceReport struct {
	Dom  float64
	Intl float64
	Bond float64
}

type RebalReport map[string]float64

func NewRebalanceReport(targetAlloc AssetAllocation, targetTotal float64, curHoldings Holdingss) RebalReport {
	targetHoldings := Holdingss{}
	for k, v := range targetAlloc {
		targetHoldings[k] = targetTotal * v
	}

	ret := RebalReport{}
	for k, v := range targetHoldings {
		ret[k] = v - curHoldings[k]
	}
	return ret
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
		topupTotal = curHoldings.dom / targetAlloc.dom
	case 1:
		topupTotal = curHoldings.intl / targetAlloc.intl
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
