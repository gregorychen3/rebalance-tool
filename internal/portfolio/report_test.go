package portfolio

import (
	"reflect"
	"testing"
)

func TestNewRebalanceReport(t *testing.T) {
	type args struct {
		targetAlloc AssetAlloc
		targetTotal float64
		curHoldings Holdings
	}
	tests := []struct {
		name string
		args args
		want RebalanceReport
	}{
		{
			name: "Rebalance report for balanced portfolio and same total should be correct",
			args: args{
				targetAlloc: AssetAlloc{
					"dom":  .6,
					"intl": .3,
					"bond": .1,
				},
				targetTotal: 100,
				curHoldings: Holdings{
					"dom":  60,
					"intl": 30,
					"bond": 10,
				},
			},
			want: RebalanceReport{
				"dom":  0,
				"intl": 0,
				"bond": 0,
			},
		},
		{
			name: "Rebalance report for unbalanced portfolio and same total should be correct",
			args: args{
				targetAlloc: AssetAlloc{
					"dom":  .6,
					"intl": .3,
					"bond": .1,
				},
				targetTotal: 10,
				curHoldings: Holdings{
					"dom":  3,
					"intl": 3,
					"bond": 4,
				},
			},
			want: RebalanceReport{
				"dom":  3,
				"intl": 0,
				"bond": -3,
			},
		},
		{
			name: "Rebalance report for balanced portfolio and different total should be correct",
			args: args{
				targetAlloc: AssetAlloc{
					"dom":  .6,
					"intl": .3,
					"bond": .1,
				},
				targetTotal: 200,
				curHoldings: Holdings{
					"dom":  60,
					"intl": 30,
					"bond": 10,
				},
			},
			want: RebalanceReport{
				"dom":  60,
				"intl": 30,
				"bond": 10,
			},
		},
		{
			name: "Rebalance report for unbalanced portfolio and different total should be correct",
			args: args{
				targetAlloc: AssetAlloc{
					"dom":  .6,
					"intl": .3,
					"bond": .1,
				},
				targetTotal: 20,
				curHoldings: Holdings{
					"dom":  3,
					"intl": 3,
					"bond": 4,
				},
			},
			want: RebalanceReport{
				"dom":  9,
				"intl": 3,
				"bond": -2,
			},
		},
		{
			name: "Rebalance report for rebalancing to total of zero should be correct",
			args: args{
				targetAlloc: AssetAlloc{
					"dom":  .6,
					"intl": .3,
					"bond": .1,
				},
				targetTotal: 0,
				curHoldings: Holdings{
					"dom":  3,
					"intl": 3,
					"bond": 4,
				},
			},
			want: RebalanceReport{
				"dom":  -3,
				"intl": -3,
				"bond": -4,
			},
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

func TestTopupTotal(t *testing.T) {
	type args struct {
		targetAlloc AssetAlloc
		curHoldings Holdings
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Topup total for a balanced portfolio should be same as its current total",
			args: args{
				targetAlloc: AssetAlloc{
					"dom":  .6,
					"intl": .3,
					"bond": .1,
				},
				curHoldings: Holdings{
					"dom":  .6,
					"intl": .3,
					"bond": .1,
				},
			},
			want: 1,
		},
		{
			name: "Topup total for a unbalanced portfolio should be correct",
			args: args{
				targetAlloc: AssetAlloc{
					"dom":  .6,
					"intl": .3,
					"bond": .1,
				},
				curHoldings: Holdings{
					"dom":  2,
					"intl": 4,
					"bond": 4,
				},
			},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopupTotal(tt.args.targetAlloc, tt.args.curHoldings); got != tt.want {
				t.Errorf("TopupTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
