package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxPoints2503(t *testing.T) {
	type args struct {
		grid    [][]int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case 1",
			args{
				[][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}},
				[]int{5, 6, 2},
			},
			[]int{5, 8, 1},
		},
		{
			"case 2",
			args{
				[][]int{{5, 2, 1}, {1, 1, 2}},
				[]int{3},
			},
			[]int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxPoints2503(tt.args.grid, tt.args.queries), "maxPoints2503(%v, %v)", tt.args.grid, tt.args.queries)
		})
	}
}
