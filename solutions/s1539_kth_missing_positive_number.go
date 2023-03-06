/*
	https://leetcode.com/problems/kth-missing-positive-number

	Given an array arr of positive integers sorted in a strictly increasing order,
	and an integer k.

	Return the kth positive integer that is missing from this array.
*/

package solutions

func findKthPositive(arr []int, k int) int {
	lo, hi := 0, len(arr)-1
	var mid int
	for lo <= hi {
		mid = (lo + hi) >> 1
		if arr[mid]-mid-1 < k {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo + k
}
