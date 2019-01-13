package portfolio

import (
	"reflect"
	"testing"
)

func TestNewAssetAlloc(t *testing.T) {
	type args struct {
		weights map[string]int
	}
	tests := []struct {
		name    string
		args    args
		want    AssetAlloc
		wantErr bool
	}{
		{
			name: "NewAssetAlloc() should return AssetAlloc with correct fields",
			args: args{
				weights: map[string]int{
					"dom":  60,
					"intl": 30,
					"bond": 10,
				},
			},
			want: AssetAlloc{
				"dom":  .6,
				"intl": .3,
				"bond": .1,
			},
			wantErr: false,
		},
		{
			name: "NewAssetAlloc() should return AssetAlloc with correct fields",
			args: args{
				weights: map[string]int{
					"dom":  60,
					"intl": 30,
					"bond": 11,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAssetAlloc(tt.args.weights)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAssetAlloc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAssetAlloc() = %v, want %v", got, tt.want)
			}
		})
	}
}
