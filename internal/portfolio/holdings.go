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

type Holdingss map[string]float64

func (h *Holdingss) Total() float64 {
	sum := 0.0
	for _, v := range *h {
		sum += v
	}
	return sum
}
