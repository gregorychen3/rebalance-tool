package portfolio

type TargetAlloc struct {
	dom  float32
	intl float32
	bond float32
}

func NewTargetAlloc(dom float32, intl float32, bond float32) *TargetAlloc {
	return &TargetAlloc{
		dom:  dom / 100,
		intl: intl / 100,
		bond: bond / 100,
	}
}
