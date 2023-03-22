/*
	https://leetcode.com/problems/range-sum-query-mutable/

	Given an integer array nums, handle multiple queries of the following types:

		Update the value of an element in nums.
		Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.

	Implement the NumArray class:

		NumArray(int[] nums) Initializes the object with the integer array nums.
		void update(int index, int val) Updates the value of nums[index] to be val.
		int sumRange(int left, int right) Returns the sum of the elements of nums between indices left
	and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).
*/

//nolint:revive // it's ok
package solutions

type NumArray struct {
	xs      []int
	prefSum []int
}

// NewNumArray should call Constructor to pass Leetcode tests
func NewNumArray(nums []int) NumArray {
	pref := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		pref[i+1] = pref[i] + nums[i]
	}
	return NumArray{xs: nums, prefSum: pref}
}

func (a *NumArray) Update(index, val int) {
	a.xs[index] = val
	for i := index; i < len(a.xs); i++ {
		a.prefSum[i+1] = a.prefSum[i] + a.xs[i]
	}
}

func (a *NumArray) SumRange(left, right int) int {
	return a.prefSum[right+1] - a.prefSum[left]
}
