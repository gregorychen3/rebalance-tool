package rebalance

type Holdings struct {
	dom  float32
	intl float32
	bond float32
}

func NewHoldings(dom float32, intl float32, bond float32) *Holdings {
	return &Holdings{
		dom:  dom,
		intl: intl,
		bond: bond,
	}
}
