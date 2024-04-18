package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_islandPerimeter(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{[][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			}},
			16,
		},
		{
			"test 2",
			args{[][]int{
				{1},
			}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, islandPerimeter(tt.args.grid), "islandPerimeter(%v)", tt.args.grid)
		})
	}
}
