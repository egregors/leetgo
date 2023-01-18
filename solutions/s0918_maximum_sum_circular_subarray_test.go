package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSubarraySumCircular(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{
				[]int{1, -2, 3, -2},
			},
			3,
		},
		{
			"test-2",
			args{
				[]int{5, -3, 5},
			},
			10,
		},
		{
			"test-3",
			args{
				[]int{-3, -2, -3},
			},
			-2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxSubarraySumCircular(tt.args.nums), "maxSubarraySumCircular(%v)", tt.args.nums)
		})
	}
}
