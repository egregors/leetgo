package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 0",
			args{obstacleGrid: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, uniquePathsWithObstacles(tt.args.obstacleGrid), "uniquePathsWithObstacles(%v)", tt.args.obstacleGrid)
		})
	}
}
