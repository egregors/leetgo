package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumUnits(t *testing.T) {
	type args struct {
		boxTypes  [][]int
		truckSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{
				boxTypes: [][]int{
					{1, 3}, {2, 2}, {3, 1},
				},
				truckSize: 4,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumUnits(tt.args.boxTypes, tt.args.truckSize), "maximumUnits(%v, %v)", tt.args.boxTypes, tt.args.truckSize)
		})
	}
}
