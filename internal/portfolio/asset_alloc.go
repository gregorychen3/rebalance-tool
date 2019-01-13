package portfolio

import "errors"

type AssetAlloc struct {
	dom  float64
	intl float64
	bond float64
}

func NewAssetAlloc(dom float64, intl float64, bond float64) *AssetAlloc {
	return &AssetAlloc{
		dom:  dom,
		intl: intl,
		bond: bond,
	}
}

type AssetAllocation map[string]int

func NewAssetAllocation(weights map[string]int) (AssetAllocation, error) {
	sum := 0
	for _, v := range weights {
		sum += v
	}
	if sum != 100 {
		return nil, errors.New("weights do not add up to 100%")
	}

	return weights, nil
}
