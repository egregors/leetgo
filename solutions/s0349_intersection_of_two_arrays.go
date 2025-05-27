/*
	https://leetcode.com/problems/intersection-of-two-arrays/description/

	Given two integer arrays nums1 and nums2, return an array of their . Each
		element in the result must
	be unique and you may return the result in any order.
*/

package solutions

func intersection(nums1 []int, nums2 []int) []int {
	m1, m2 := make(map[int]struct{}), make(map[int]struct{})
	for _, n := range nums1 {
		m1[n] = struct{}{}
	}
	for _, n := range nums2 {
		m2[n] = struct{}{}
	}

	var res []int

	for k := range m1 {
		if _, ok := m2[k]; ok {
			res = append(res, k)
		}
	}

	return res
}
