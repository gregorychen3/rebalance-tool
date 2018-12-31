package rebalance

type TargetAlloc struct {
	dom  float32
	intl float32
	bond float32
}

func NewTargetAlloc(dom float32, intl float32, bond float32) *TargetAlloc {
	return &TargetAlloc{
		dom:  dom,
		intl: intl,
		bond: bond,
	}
}
