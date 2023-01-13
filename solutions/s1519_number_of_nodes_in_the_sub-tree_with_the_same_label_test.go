package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countSubTrees(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		labels string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{
				7,
				[][]int{
					{0, 1},
					{0, 2},
					{1, 4},
					{1, 5},
					{2, 3},
					{2, 6},
				},
				"abaedcd",
			},
			[]int{2, 1, 1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countSubTrees(tt.args.n, tt.args.edges, tt.args.labels), "countSubTrees(%v, %v, %v)", tt.args.n, tt.args.edges, tt.args.labels)
		})
	}
}
