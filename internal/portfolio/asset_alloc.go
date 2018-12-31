package portfolio

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
