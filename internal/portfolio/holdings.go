package portfolio

type Holdings struct {
	dom  float64
	intl float64
	bond float64
}

func NewHoldings(dom float64, intl float64, bond float64) *Holdings {
	return &Holdings{
		dom:  dom,
		intl: intl,
		bond: bond,
	}
}

func (h *Holdings) Total() float64 {
	return h.dom + h.intl + h.bond
}
