/*
	https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/description/

	Given the array of integers nums, you will choose two different indices i and j of that array.
	Return the maximum value of (nums[i]-1)*(nums[j]-1).
*/

package solutions

import "testing"

func Test_maxProduct1464(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{nums: []int{3, 4, 5, 2}},
			want: 12,
		},
		{
			name: "Test 2",
			args: args{nums: []int{1, 5, 4, 5}},
			want: 16,
		},
		{
			name: "Test 3",
			args: args{nums: []int{3, 7}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct1464(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct1464() = %v, want %v", got, tt.want)
			}
		})
	}
}
