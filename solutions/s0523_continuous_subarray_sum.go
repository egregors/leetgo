/*
	https://leetcode.com/problems/continuous-subarray-sum/

	Given an integer array nums and an integer k, return true if nums has a continuous subarray of size
	at least two whose elements sum up to a multiple of k, or false otherwise.

	An integer x is a multiple of k if there exists an integer n such that x = n * k. 0 is always a multiple of k.
*/

package solutions

func checkSubarraySum(nums []int, k int) bool {
	m := make(map[int]int)
	m[0] = 0
	s := 0
	for i, n := range nums {
		s += n
		if _, ok := m[s%k]; !ok {
			m[s%k] = i + 1
		} else if x := m[s%k]; x < i {
			return true
		}
	}
	return false
}
