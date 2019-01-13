package portfolio

type Holdings map[string]float64

func (h *Holdings) Total() float64 {
	sum := 0.0
	for _, v := range *h {
		sum += v
	}
	return sum
}
