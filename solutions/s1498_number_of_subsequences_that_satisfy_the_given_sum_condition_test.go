package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numSubseq(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{[]int{3, 5, 6, 7}, 9},
			4,
		},
		{
			"test 2",
			args{[]int{3, 3, 6, 8}, 10},
			6,
		},
		{
			"test 3",
			args{[]int{2, 3, 3, 4, 6, 7}, 12},
			61,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numSubseq(tt.args.nums, tt.args.target), "numSubseq(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}
