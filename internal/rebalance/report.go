package rebalance

type RebalanceReport struct {
	dom  float32
	intl float32
	bond float32
}

func NewRebalanceReport(target *TargetAlloc, curHoldings *Holdings) *RebalanceReport {
	total := curHoldings.total()
	targetHoldings := Holdings{
		dom:  total * target.dom,
		intl: total * target.intl,
		bond: total * target.bond,
	}
	return &RebalanceReport{
		dom:  targetHoldings.dom - curHoldings.dom,
		intl: targetHoldings.intl - curHoldings.intl,
		bond: targetHoldings.bond - curHoldings.bond,
	}
}
