package portfolio

type Holdingss map[string]float64

func (h *Holdingss) Total() float64 {
	sum := 0.0
	for _, v := range *h {
		sum += v
	}
	return sum
}
