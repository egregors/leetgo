package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumFuelCost(t *testing.T) {
	type args struct {
		roads [][]int
		seats int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"Example 1",
			args{
				[][]int{
					{0, 1},
					{0, 2},
					{0, 3},
				},
				5,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumFuelCost(tt.args.roads, tt.args.seats), "minimumFuelCost(%v, %v)", tt.args.roads, tt.args.seats)
		})
	}
}
