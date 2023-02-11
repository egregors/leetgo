package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shortestAlternatingPaths(t *testing.T) {
	type args struct {
		n         int
		redEdges  [][]int
		blueEdges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test 1",
			args{
				3,
				[][]int{{0, 1}, {1, 2}},
				[][]int{},
			},
			[]int{0, 1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, shortestAlternatingPaths(tt.args.n, tt.args.redEdges, tt.args.blueEdges), "shortestAlternatingPaths(%v, %v, %v)", tt.args.n, tt.args.redEdges, tt.args.blueEdges)
		})
	}
}
