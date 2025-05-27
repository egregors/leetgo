/*
	https://leetcode.com/problems/minimize-maximum-of-array/

	You are given a 0-indexed array nums comprising of n non-negative integers.

	In one operation, you must:

		Choose an integer i such that 1 <= i < n and nums[i] > 0.
		Decrease nums[i] by 1.
		Increase nums[i - 1] by 1.

	Return the minimum possible value of the maximum integer of nums after
		performing
	any number of operations.
*/

package solutions

func minimizeArrayValue(nums []int) int {
	var res, pref int
	for i := 0; i < len(nums); i++ {
		pref += nums[i]
		res = Maximum(res, (pref+i)/(i+1))
	}
	return res
}
