package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubsequence(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//Input: nums = [2,1,3,3], k = 2
		//Output: [3,3]
		{
			name: "test 1",
			args: args{
				nums: []int{2, 1, 3, 3},
				k:    2,
			},
			want: []int{3, 3},
		},
		//Input: nums = [-1,-2,3,4], k = 3
		//Output: [-1,3,4]
		{
			name: "test 2",
			args: args{
				nums: []int{-1, -2, 3, 4},
				k:    3,
			},
			want: []int{-1, 3, 4},
		},
		//Input: nums = [3,4,3,3], k = 2
		//Output: [3,4]
		{
			name: "test 3",
			args: args{
				nums: []int{3, 4, 3, 3},
				k:    2,
			},
			want: []int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxSubsequence(tt.args.nums, tt.args.k), "maxSubsequence(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
