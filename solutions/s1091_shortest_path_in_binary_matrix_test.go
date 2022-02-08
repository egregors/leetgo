package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shortestPathBinaryMatrix(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 0",
			args{grid: [][]int{
				{0, 0, 0, 1, 0, 0, 0},
				{0, 1, 1, 1, 0, 0, 0},
				{0, 1, 0, 0, 0, 1, 1},
				{0, 0, 1, 1, 1, 0, 1},
				{0, 1, 1, 1, 0, 0, 0},
			}},
			9,
		},
		{
			"Test 1",
			args{grid: [][]int{{0, 1}, {1, 0}}},
			2,
		},
		{
			"Test 2",
			args{grid: [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}},
			4,
		},
		{
			"Test 3",
			args{grid: [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, shortestPathBinaryMatrix(tt.args.grid), "shortestPathBinaryMatrix(%v)", tt.args.grid)
		})
	}
}
