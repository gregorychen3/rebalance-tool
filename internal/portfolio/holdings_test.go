package portfolio

import "testing"

func TestHoldings_Total(t *testing.T) {
	tests := []struct {
		name string
		h    *Holdings
		want float64
	}{
		{
			name: "Total holdings should be correct",
			h: &Holdings{
				"foo": 1,
				"bar": 2,
			},
			want: 3,
		},
		{
			name: "Total holdings of empty portfolio should be zero",
			h:    &Holdings{},
			want: 0,
		},
		{
			name: "Total holdings of portfolio with one asset should be correct",
			h: &Holdings{
				"foo": 1,
			},
			want: 1,
		},
		{
			name: "Total holdings of portfolio with many assets should be correct",
			h: &Holdings{
				"foo": 1,
				"bar": 2,
				"baz": 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Total(); got != tt.want {
				t.Errorf("Holdings.Total() = %v, want %v", got, tt.want)
			}
		})
	}
}
