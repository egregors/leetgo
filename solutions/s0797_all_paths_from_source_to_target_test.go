package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Test 1",
			args{graph: [][]int{{1, 2}, {3}, {3}, {}}},
			[][]int{{0, 1, 3}, {0, 2, 3}},
		},
		{
			"Test 2",
			args{graph: [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}},
			[][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, allPathsSourceTarget(tt.args.graph), "allPathsSourceTarget(%v)", tt.args.graph)
		})
	}
}
