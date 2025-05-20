package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isZeroArray(t *testing.T) {
	type args struct {
		nums    []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				[]int{1, 0, 1},
				[][]int{{0, 2}},
			},
			true,
		},
		{
			"test2",
			args{
				[]int{4, 3, 2, 1},
				[][]int{{1, 3}, {0, 2}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isZeroArray(tt.args.nums, tt.args.queries), "isZeroArray(%v, %v)", tt.args.nums, tt.args.queries)
		})
	}
}
