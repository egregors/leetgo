package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			3,
		},
		{
			"2",
			args{
				gas:  []int{2, 3, 4},
				cost: []int{3, 4, 3},
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canCompleteCircuit(tt.args.gas, tt.args.cost), "canCompleteCircuit(%v, %v)", tt.args.gas, tt.args.cost)
		})
	}
}
