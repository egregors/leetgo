package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//Example 1:
		//
		//Input: nums = [1,2,3,4,5]
		//
		//Output: 15
		//
		//Explanation:
		//
		//Select the entire array without deleting any element to obtain the maximum sum.
		//
		//Example 2:
		//
		//Input: nums = [1,1,0,1,1]
		//
		//Output: 1
		//
		//Explanation:
		//
		//Delete the element nums[0] == 1, nums[1] == 1, nums[2] == 0, and nums[3] == 1. Select the entire array [1] to obtain the maximum sum.
		//
		//Example 3:
		//
		//Input: nums = [1,2,-1,-2,1,0,-1]
		//
		//Output: 3
		//
		//Explanation:
		//
		//Delete the elements nums[2] == -1 and nums[3] == -2, and select the subarray [2, 1] from [1, 2, 1, 0, -1] to obtain the maximum sum.
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1, 1, 0, 1, 1},
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1, 2, -1, -2, 1, 0, -1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxSum(tt.args.nums), "maxSum(%v)", tt.args.nums)
		})
	}
}
