package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minTime(t *testing.T) {
	type args struct {
		n        int
		edges    [][]int
		hasApple []bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
				[]bool{false, false, true, false, true, true, false},
			},
			8,
		},
		{
			"test 2",
			args{
				7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
				[]bool{false, false, true, false, false, true, false},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minTime(tt.args.n, tt.args.edges, tt.args.hasApple), "minTime(%v, %v, %v)", tt.args.n, tt.args.edges, tt.args.hasApple)
		})
	}
}
