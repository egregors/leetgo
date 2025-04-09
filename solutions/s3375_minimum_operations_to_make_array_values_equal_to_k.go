/*
	https://leetcode.com/problems/minimum-operations-to-make-array-values-equal-to-k/description/

	You are given an integer array nums and an integer k.

	An integer h is called valid if all values in the array that are strictly greater than h
	are identical.

	For example, if nums = [10, 8, 10, 8], a valid integer is h = 9 because all nums[i] > 9
	are equal to 10, but 5 is not a valid integer.

	You are allowed to perform the following operation on nums:

		Select an integer h that is valid for the current values in nums.
		For each index i where nums[i] > h, set nums[i] to h.

	Return the minimum number of operations required to make every element in nums equal to k.
	If it is impossible to make all elements equal to k, return -1.
*/

package solutions

func minOperations3375(nums []int, k int) int {
	m := make(map[int]struct{})
	for _, n := range nums {
		if n < k {
			return -1
		}
		if n > k {
			m[n] = struct{}{}
		}
	}

	return len(m)
}
