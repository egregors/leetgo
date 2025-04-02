/*
	https://leetcode.com/problems/find-the-integer-added-to-array-i/description/

	You are given two arrays of equal length, nums1 and nums2.

	Each element in nums1 has been increased (or decreased in the case of negative)
	by an integer, represented by the variable x.

	As a result, nums1 becomes equal to nums2. Two arrays are considered equal when
	they contain the same integers with the same frequencies.

	Return the integer x.
*/

package solutions

import "testing"

func Test_addedInteger(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums1: []int{2, 6, 4},
				nums2: []int{9, 7, 5},
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				nums1: []int{10},
				nums2: []int{5},
			},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addedInteger(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("addedInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
