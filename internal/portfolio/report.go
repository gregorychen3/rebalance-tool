package portfolio

type RebalanceReport struct {
	Dom  float32
	Intl float32
	Bond float32
}

func NewRebalanceReport(target *AssetAlloc, curHoldings *Holdings) *RebalanceReport {
	total := curHoldings.Total()
	targetHoldings := Holdings{
		dom:  total * target.dom,
		intl: total * target.intl,
		bond: total * target.bond,
	}
	return &RebalanceReport{
		Dom:  targetHoldings.dom - curHoldings.dom,
		Intl: targetHoldings.intl - curHoldings.intl,
		Bond: targetHoldings.bond - curHoldings.bond,
	}
}
