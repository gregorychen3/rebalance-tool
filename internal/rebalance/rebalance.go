package rebalance

type Weightings struct {
	dom  float32
	intl float32
	bond float32
}

func NewWeightings(dom float32, intl float32, bond float32) *Weightings {
	return &Weightings{
		dom:  dom,
		intl: intl,
		bond: bond,
	}
}
