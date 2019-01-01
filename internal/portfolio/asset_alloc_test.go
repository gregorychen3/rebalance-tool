package portfolio

import (
	"reflect"
	"testing"
)

func TestNewAssetAlloc(t *testing.T) {
	type args struct {
		dom  float64
		intl float64
		bond float64
	}
	tests := []struct {
		name string
		args args
		want *AssetAlloc
	}{
		{
			name: "NewAssetAlloc() should return AssetAlloc with correct fields",
			args: args{
				dom:  1.1,
				intl: 2.2,
				bond: 3.3,
			},
			want: &AssetAlloc{
				dom:  1.1,
				intl: 2.2,
				bond: 3.3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAssetAlloc(tt.args.dom, tt.args.intl, tt.args.bond); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAssetAlloc() = %v, want %v", got, tt.want)
			}
		})
	}
}
