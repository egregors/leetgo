package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				costs: [][]int{
					{17, 2, 17}, {16, 16, 5}, {14, 3, 19},
				},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minCost(tt.args.costs), "minCost(%v)", tt.args.costs)
		})
	}
}
