package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findSmallestSetOfVertices(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}},
			},
			want: []int{0, 3},
		},
		{
			name: "test 2",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {0, 2}, {3, 4}},
			},
			want: []int{0, 3},
		},
		{
			name: "test 3",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findSmallestSetOfVertices(tt.args.n, tt.args.edges), "findSmallestSetOfVertices(%v, %v)", tt.args.n, tt.args.edges)
		})
	}
}
