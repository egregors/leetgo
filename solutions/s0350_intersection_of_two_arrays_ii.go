/*
	https://leetcode.com/problems/intersection-of-two-arrays-ii/

	Given two integer arrays nums1 and nums2, return an array of their
		intersection.
	Each element in the result must appear as many times as it shows in both arrays
		and
	you may return the result in any order.
*/

package solutions

func intersect(nums1, nums2 []int) []int {
	m := make(map[int]int)
	short := nums2
	long := nums1

	if len(nums1) < len(nums2) {
		short, long = nums1, nums2

	}

	for _, v := range short {
		m[v]++
	}

	var res []int
	for _, v := range long {
		if m[v] > 0 {
			res = append(res, v)
			m[v]--
		}
	}

	return res
}
