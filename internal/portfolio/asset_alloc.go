package portfolio

import "errors"

type AssetAllocation map[string]float64

func NewAssetAllocation(weights map[string]int) (AssetAllocation, error) {
	ret := AssetAllocation{}
	sum := 0
	for k, v := range weights {
		ret[k] = float64(v) / 100.0
		sum += v
	}
	if sum != 100 {
		return nil, errors.New("weights do not add up to 100%")
	}

	return ret, nil
}
