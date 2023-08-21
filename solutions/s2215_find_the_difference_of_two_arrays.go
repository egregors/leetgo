/*
	https://leetcode.com/problems/find-the-difference-of-two-arrays/

		Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

		answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
		answer[1] is a list of all distinct integers in nums2 which are not present in nums1.

	Note that the integers in the lists may be returned in any order.
*/

package solutions

func findDifference(nums1, nums2 []int) [][]int {
	m1, m2 := make(map[int]struct{}, len(nums1)), make(map[int]struct{}, len(nums2))

	for _, n := range nums1 {
		m1[n] = struct{}{}
	}
	for _, n := range nums2 {
		m2[n] = struct{}{}
	}

	var res1, res2 []int

	for k := range m1 {
		if _, ok := m2[k]; !ok {
			res1 = append(res1, k)
		}
	}
	for k := range m2 {
		if _, ok := m1[k]; !ok {
			res2 = append(res2, k)
		}
	}

	return [][]int{res1, res2}
}
