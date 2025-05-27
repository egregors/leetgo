/*
	https://leetcode.com/problems/maximum-erasure-value/

	You are given an array of positive integers nums and want to erase a subarray
		containing
	unique elements. The score you get by erasing the subarray is equal to the sum
		of its elements.

	Return the maximum score you can get by erasing exactly one subarray.

	An array b is called to be a subarray of a if it forms a contiguous subsequence
		of a,
	that is, if it is equal to a[l],a[l+1],...,a[r] for some (l,r).
*/

package solutions

func maximumUniqueSubarray(nums []int) int {
	res, curr, start := 0, 0, 0
	set := make(Set[int])
	for end := 0; end < len(nums); end++ {
		for set.Contains(nums[end]) {
			set.Remove(nums[start])
			curr -= nums[start]
			start++
		}
		curr += nums[end]
		set.Add(nums[end])
		res = Maximum(res, curr)
	}
	return res
}
