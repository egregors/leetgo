package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countPairs(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test 1",
			args: args{
				n:     3,
				edges: [][]int{{0, 1}, {0, 2}, {1, 2}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countPairs(tt.args.n, tt.args.edges), "countPairs(%v, %v)", tt.args.n, tt.args.edges)
		})
	}
}
