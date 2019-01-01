package portfolio

import (
	"math"
	"reflect"
	"testing"
)

func TestTopupTotal(t *testing.T) {
	FLOAT_COMPARE_THRESHOLD := .01
	type args struct {
		targetAlloc *AssetAlloc
		curHoldings *Holdings
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Topup total for a balanced portfolio should be the portfolio's total value",
			args: args{
				targetAlloc: &AssetAlloc{dom: .6, intl: .3, bond: .1},
				curHoldings: &Holdings{dom: 60, intl: 30, bond: 10},
			},
			want: 100,
		},
		{
			name: "Topup total for a drifted portfolio should be correct",
			args: args{
				targetAlloc: &AssetAlloc{dom: .6, intl: .3, bond: .1},
				curHoldings: &Holdings{dom: 33.33, intl: 33.33, bond: 33.33},
			},
			want: 333.3,
		},
		{
			name: "Topup total for an empty portfolio should be 0",
			args: args{
				targetAlloc: &AssetAlloc{dom: .6, intl: .3, bond: .1},
				curHoldings: &Holdings{dom: 0, intl: 0, bond: 0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopupTotal(tt.args.targetAlloc, tt.args.curHoldings); math.Abs(got-tt.want) > FLOAT_COMPARE_THRESHOLD {
				t.Errorf("TopupTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRebalanceReport(t *testing.T) {
	type args struct {
		targetAlloc *AssetAlloc
		targetTotal float64
		curHoldings *Holdings
	}
	tests := []struct {
		name string
		args args
		want *RebalanceReport
	}{
		{
			name: "Rebalance report for an empty portfolio should be correct",
			args: args{
				targetAlloc: &AssetAlloc{dom: .6, intl: .3, bond: .1},
				targetTotal: 0,
				curHoldings: &Holdings{dom: 0, intl: 0, bond: 0},
			},
			want: &RebalanceReport{Dom: 0, Intl: 0, Bond: 0},
		},
		{
			name: "Rebalance report for a balanced portfolio should be correct",
			args: args{
				targetAlloc: &AssetAlloc{dom: .6, intl: .3, bond: .1},
				targetTotal: 100,
				curHoldings: &Holdings{dom: 60, intl: 30, bond: 10},
			},
			want: &RebalanceReport{Dom: 0, Intl: 0, Bond: 0},
		},
		{
			name: "Rebalance report for a drifted portfolio should be correct",
			args: args{
				targetAlloc: &AssetAlloc{dom: .6, intl: .3, bond: .1},
				targetTotal: 100,
				curHoldings: &Holdings{dom: 50, intl: 40, bond: 10},
			},
			want: &RebalanceReport{Dom: 10, Intl: -10, Bond: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRebalanceReport(tt.args.targetAlloc, tt.args.targetTotal, tt.args.curHoldings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRebalanceReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
