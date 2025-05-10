/*
	https://leetcode.com/problems/minimum-equal-sum-of-two-arrays-after-replacing-zeros/description/

	You are given two arrays nums1 and nums2 consisting of positive integers.

	You have to replace all the 0's in both arrays with strictly positive integers such that the sum of elements
	of both arrays becomes equal.

	Return the minimum equal sum you can obtain, or -1 if it is impossible.
*/

package solutions

func minSum(nums1 []int, nums2 []int) int64 {
	var sum1, zs1, sum2, zs2 int
	for _, n := range nums1 {
		sum1 += n
		if n == 0 {
			sum1++
			zs1++
		}
	}

	for _, n := range nums2 {
		sum2 += n
		if n == 0 {
			sum2++
			zs2++
		}
	}

	if (zs1 == 0 && sum2 > sum1) || (zs2 == 0 && sum1 > sum2) {
		return -1
	}

	return int64(max(sum1, sum2))
}
