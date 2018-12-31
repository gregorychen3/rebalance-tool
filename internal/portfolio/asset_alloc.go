package portfolio

type AssetAlloc struct {
	dom  float32
	intl float32
	bond float32
}

func NewAssetAlloc(dom float32, intl float32, bond float32) *AssetAlloc {
	return &AssetAlloc{
		dom:  dom / 100,
		intl: intl / 100,
		bond: bond / 100,
	}
}
