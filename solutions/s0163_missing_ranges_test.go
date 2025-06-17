package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMissingRanges(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				[]int{0, 1, 3, 50, 75},
				0,
				99,
			},
			[][]int{{2, 2}, {4, 49}, {51, 74}, {76, 99}},
		},
		{
			"2",
			args{
				[]int{-1},
				-1,
				-1,
			},
			[][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMissingRanges(tt.args.nums, tt.args.lower, tt.args.upper), "findMissingRanges(%v, %v, %v)", tt.args.nums, tt.args.lower, tt.args.upper)
		})
	}
}
