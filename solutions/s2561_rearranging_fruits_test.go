package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCost2561(t *testing.T) {
	type args struct {
		basket1 []int
		basket2 []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				basket1: []int{4, 2, 2, 2},
				basket2: []int{1, 4, 1, 2},
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				basket1: []int{2, 3, 4, 1},
				basket2: []int{3, 2, 5, 1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minCost2561(tt.args.basket1, tt.args.basket2), "minCost2561(%v, %v)", tt.args.basket1, tt.args.basket2)
		})
	}
}
