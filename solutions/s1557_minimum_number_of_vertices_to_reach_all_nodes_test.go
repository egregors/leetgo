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
		want map[int]bool
	}{
		{
			name: "test 1",
			args: args{
				n:     6,
				edges: [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}},
			},
			want: map[int]bool{0: true, 3: true},
		},
		{
			name: "test 2",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {0, 2}, {3, 4}},
			},
			want: map[int]bool{0: true, 3: true},
		},
		{
			name: "test 3",
			args: args{
				n:     5,
				edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			},
			want: map[int]bool{0: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findSmallestSetOfVertices(tt.args.n, tt.args.edges)
			resM := make(map[int]bool)
			for _, n := range res {
				resM[n] = true
			}
			for k := range resM {
				if _, ok := tt.want[k]; !ok {
					assert.Equalf(t, tt.want[k], k, "findSmallestSetOfVertices(%v, %v)", tt.args.n, tt.args.edges)
				}
			}
		})
	}
}
