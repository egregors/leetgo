/*
	https://leetcode.com/problems/subarray-product-less-than-k/

	Given an array of integers nums and an integer k, return the number of
		contiguous subarrays
	where the product of all the elements in the subarray is strictly less than k.
*/

package solutions

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	prod, res, l := 1, 0, 0
	for r := 0; r < len(nums); r++ {
		prod *= nums[r]
		for prod >= k {
			prod /= nums[l]
			l++
		}
		res += r - l + 1
	}
	return res
}
