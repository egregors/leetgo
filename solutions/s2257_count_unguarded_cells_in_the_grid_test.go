package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countUnguarded(t *testing.T) {
	type args struct {
		m      int
		n      int
		guards [][]int
		walls  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//Input: m = 4, n = 6, guards = [[0,0],[1,1],[2,3]], walls = [[0,1],[2,2],[1,4]]
		//Output: 7
		{
			name: "test case 1",
			args: args{
				m:      4,
				n:      6,
				guards: [][]int{{0, 0}, {1, 1}, {2, 3}},
				walls:  [][]int{{0, 1}, {2, 2}, {1, 4}},
			},
			want: 7,
		},
		//Input: m = 3, n = 3, guards = [[1,1]], walls = [[0,1],[1,0],[2,1],[1,2]]
		//Output: 4
		{
			name: "test case 2",
			args: args{
				m:      3,
				n:      3,
				guards: [][]int{{1, 1}},
				walls:  [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countUnguarded(tt.args.m, tt.args.n, tt.args.guards, tt.args.walls), "countUnguarded(%v, %v, %v, %v)", tt.args.m, tt.args.n, tt.args.guards, tt.args.walls)
		})
	}
}
