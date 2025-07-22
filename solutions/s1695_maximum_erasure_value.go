/*
	https://leetcode.com/problems/maximum-erasure-value/

	You are given an array of positive integers nums and want to erase a subarray
		containing unique elements. The score you get by erasing the subarray is equal
		to the sum of its elements.

	Return the maximum score you can get by erasing exactly one subarray.

	An array b is called to be a subarray of a if it forms a contiguous subsequence
		of a, that is, if it is equal to a[l],a[l+1],...,a[r] for some (l,r).
*/

package solutions

func maximumUniqueSubarray1695(nums []int) int {
	var l, r, res, sum int
	m := make([]int, 10_001)
	for ; r < len(nums); r++ {
		el := nums[r]
		m[el]++
		sum += el
		for l < r && m[el] > 1 {
			m[nums[l]]--
			sum -= nums[l]
			l++
		}
		res = max(res, sum)
	}

	return res
}
