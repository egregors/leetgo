package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minZeroArray(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				[]int{2, 0, 2},
				[][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}},
			},
			2,
		},
		{
			"case 2",
			args{
				[]int{4, 3, 2, 1},
				[][]int{{1, 3, 2}, {0, 2, 1}},
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minZeroArray(tt.args.nums, tt.args.queries), "minZeroArray(%v, %v)", tt.args.nums, tt.args.queries)
		})
	}
}
