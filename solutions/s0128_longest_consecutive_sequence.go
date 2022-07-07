/*
	https://leetcode.com/problems/longest-consecutive-sequence/

	Given an unsorted array of integers nums, return the length of the longest
	consecutive elements sequence.

	You must write an algorithm that runs in O(n) time.
*/

package solutions

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		m[n] = struct{}{}
	}

	long := 0
	for _, n := range nums {
		if _, ok := m[n-1]; !ok {
			count := 0

			_, ok := m[n]
			for ok {
				count++
				n++
				_, ok = m[n]
			}

			if count > long {
				long = count
			}
		}
	}
	return long
}
