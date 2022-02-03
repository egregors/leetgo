package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numSubarrayProductLessThanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{10, 5, 2, 6},
				k:    100,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numSubarrayProductLessThanK(tt.args.nums, tt.args.k), "numSubarrayProductLessThanK(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
