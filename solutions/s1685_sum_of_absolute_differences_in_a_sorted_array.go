/*
	https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/description/

	You are given an integer array nums sorted in non-decreasing order.

	Build and return an integer array result with the same length as nums such that result[i] is equal
	to the summation of absolute differences between nums[i] and all the other elements in the array.

	In other words, result[i] is equal to sum(|nums[i]-nums[j]|) where 0 <= j < nums.length and j != i (0-indexed).
*/

package solutions

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	pref, suff := 0, Sum(nums...)
	res := make([]int, n)

	for i, num := range nums {
		res[i] = suff - pref - num*(n-i-i)
		suff -= num
		pref += num
	}

	return res
}
