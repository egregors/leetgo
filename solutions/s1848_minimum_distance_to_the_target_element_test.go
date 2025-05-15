/*
	https://leetcode.com/problems/minimum-distance-to-the-target-element

	Given an integer array nums (0-indexed) and two integers target and start, find an index i such that
	nums[i] == target and abs(i - start) is minimized. Note that abs(x) is the absolute value of x.

	Return abs(i - start).

	It is guaranteed that target exists in nums.
*/

package solutions

import "testing"

func Test_getMinDistance(t *testing.T) {
	type args struct {
		nums   []int
		target int
		start  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 5,
				start:  3,
			},
			1,
		},
		{
			"test2",
			args{
				nums:   []int{1},
				target: 1,
				start:  0,
			},
			0,
		},
		{
			"test3",
			args{
				[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				1,
				0,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinDistance(tt.args.nums, tt.args.target, tt.args.start); got != tt.want {
				t.Errorf("getMinDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
