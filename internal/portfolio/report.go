package portfolio

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

func TopupTotal(targetAlloc AssetAllocation, curHoldings Holdingss) float64 {
	// find most overweighted category
	curTotal := curHoldings.Total()
	curAlloc := AssetAllocation{}
	for k, v := range curHoldings {
		curAlloc[k] = v / curTotal
	}

	allocDiffs := map[string]float64{}
	for k, v := range curAlloc {
		allocDiffs[k] = v - targetAlloc[k]
	}

	mostOverallocatedAsset := getKeyWithGreatestValue(allocDiffs)
	return curHoldings[mostOverallocatedAsset] / targetAlloc[mostOverallocatedAsset]
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
