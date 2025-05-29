/*
	https://leetcode.com/problems/next-greater-element-i/

	The next greater element of some element x in an array is the first greater
		element that is to the right of x in the same array.

	You are given two distinct 0-indexed integer arrays nums1 and nums2, where
		nums1 is a subset of nums2.

	For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j]
		and determine the next greater element of nums2[j] in nums2. If there is no
		next greater element, then the answer for this query is -1.

	Return an array ans of length nums1.length such that ans[i] is the next greater
		element as described above.
*/

package solutions

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	s := []int{}
	m := make(map[int]int)

	for i := 0; i < len(nums2); i++ {
		for len(s) > 0 && nums2[i] > s[len(s)-1] {
			m[s[len(s)-1]] = nums2[i]
			s = s[:len(s)-1]
		}
		s = append(s, nums2[i])
	}

	for _, n := range s {
		m[n] = -1
	}

	res := make([]int, 0, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res = append(res, m[nums1[i]])
	}

	return res
}
