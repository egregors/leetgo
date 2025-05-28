package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxTargetNodes(t *testing.T) {
	type args struct {
		edges1 [][]int
		edges2 [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//Input: edges1 = [[0,1],[0,2],[2,3],[2,4]], edges2 = [[0,1],[0,2],[0,3],[2,7],[1,4],[4,5],[4,6]], k = 2
		//
		//Output: [9,7,9,8,8]
		{
			name: "Test 1",
			args: args{
				edges1: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}},
				edges2: [][]int{{0, 1}, {0, 2}, {0, 3}, {2, 7}, {1, 4}, {4, 5}, {4, 6}},
				k:      2,
			},
			want: []int{9, 7, 9, 8, 8},
		},
		// Input: edges1 = [[0,1],[0,2],[0,3],[0,4]], edges2 = [[0,1],[1,2],[2,3]], k = 1
		//
		//Output: [6,3,3,3,3]
		{
			name: "Test 2",
			args: args{
				edges1: [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}},
				edges2: [][]int{{0, 1}, {1, 2}, {2, 3}},
				k:      1,
			},
			want: []int{6, 3, 3, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxTargetNodes(tt.args.edges1, tt.args.edges2, tt.args.k), "maxTargetNodes(%v, %v, %v)", tt.args.edges1, tt.args.edges2, tt.args.k)
		})
	}
}
