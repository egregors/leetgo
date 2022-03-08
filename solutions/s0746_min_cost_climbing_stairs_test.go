package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				cost: []int{10, 15, 20},
			},
			want: 15,
		},
		{
			name: "Test 2",
			args: args{
				cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minCostClimbingStairs(tt.args.cost), "minCostClimbingStairs(%v)", tt.args.cost)
		})
	}
}
