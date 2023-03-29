package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minPathSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example 1",
			args{
				[][]int{
					{1, 3, 1},
					{1, 5, 1},
					{4, 2, 1},
				}},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minPathSum(tt.args.grid), "minPathSum(%v)", tt.args.grid)
		})
	}
}
