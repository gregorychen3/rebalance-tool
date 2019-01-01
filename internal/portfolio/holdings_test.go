package portfolio

import (
	"reflect"
	"testing"
)

func TestNewHoldings(t *testing.T) {
	type args struct {
		dom  float64
		intl float64
		bond float64
	}
	tests := []struct {
		name string
		args args
		want *Holdings
	}{
		{
			name: "NewHoldings() should return Holdings with correct fields",
			args: args{
				dom:  1.1,
				intl: 2.2,
				bond: 3.3,
			},
			want: &Holdings{
				dom:  1.1,
				intl: 2.2,
				bond: 3.3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHoldings(tt.args.dom, tt.args.intl, tt.args.bond); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHoldings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHoldings_Total(t *testing.T) {
	type fields struct {
		dom  float64
		intl float64
		bond float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Empty Holdings should have total of 0",
			fields: fields{
				dom:  0,
				intl: 0,
				bond: 0,
			},
			want: 0,
		},
		{
			name: "Non-empty Holdings should have correct total",
			fields: fields{
				dom:  1.1,
				intl: 2.2,
				bond: 3.3,
			},
			want: 6.6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Holdings{
				dom:  tt.fields.dom,
				intl: tt.fields.intl,
				bond: tt.fields.bond,
			}
			if got := h.Total(); got != tt.want {
				t.Errorf("Holdings.Total() = %v, want %v", got, tt.want)
			}
		})
	}
}
