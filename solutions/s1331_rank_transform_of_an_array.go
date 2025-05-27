/*
	https://leetcode.com/problems/rank-transform-of-an-array/

		Given an array of integers arr, replace each element with its rank.

	The rank represents how large the element is. The rank has the following rules:

		Rank is an integer starting from 1.
		The larger the element, the larger the rank. If two elements are equal, their
			rank must be the same.
		Rank should be as small as possible.
*/

package solutions

import "sort"

func arrayRankTransform(arr []int) []int {
	cp := make([]int, len(arr))
	copy(cp, arr)

	sort.Ints(arr)
	m := make(map[int]int)
	r := 1
	for _, n := range arr {
		if _, ok := m[n]; !ok {
			m[n] = r
			r++
		}
	}

	for i := 0; i < len(cp); i++ {
		cp[i] = m[cp[i]]
	}

	return cp
}
