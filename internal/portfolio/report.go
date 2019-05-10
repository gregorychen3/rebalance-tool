package portfolio

type RebalanceReport map[string]float64

func NewRebalanceReport(targetAlloc AssetAlloc, targetTotal float64, curHoldings Holdings) RebalanceReport {
	targetHoldings := Holdings{}
	for k, v := range targetAlloc {
		targetHoldings[k] = targetTotal * v
	}

	ret := RebalanceReport{}
	for k, v := range targetHoldings {
		ret[k] = v - curHoldings[k]
	}
	return ret
}

func TopupTotal(targetAlloc AssetAlloc, curHoldings Holdings) float64 {
	// find most overweighted category
	curTotal := curHoldings.Total()
	curAlloc := AssetAlloc{}
	for k, v := range curHoldings {
		curAlloc[k] = v / curTotal
	}

	allocDiffs := map[string]float64{}
	for k, v := range curAlloc {
		allocDiffs[k] = (v - targetAlloc[k]) / targetAlloc[k]
	}

	mostOverallocatedAsset := getKeyWithGreatestValue(allocDiffs)
	return curHoldings[mostOverallocatedAsset] / targetAlloc[mostOverallocatedAsset]
}

func getKeyWithGreatestValue(m map[string]float64) string {
	var maxK string
	var maxV float64
	for k, v := range m {
		maxK = k
		maxV = v
		break
	}
	for k, v := range m {
		if v > maxV {
			maxK = k
			maxV = v
		}
	}
	return maxK
}
