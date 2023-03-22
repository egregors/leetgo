/*
	https://leetcode.com/problems/reduce-array-size-to-the-half/

		You are given an integer array arr. You can choose a set of integers and remove all the
	occurrences of these integers in the array.

	Return the minimum size of the set so that at least half of the integers of the array are removed.
*/

package solutions

import "sort"

func minSetSize(arr []int) int {
	m := make(map[int]int)
	for _, n := range arr {
		m[n]++
	}

	var xs [][]int //nolint:prealloc // meh
	for k, v := range m {
		xs = append(xs, []int{k, v})
	}

	sort.Slice(xs, func(i, j int) bool {
		return xs[i][1] > xs[j][1]
	})

	arrLen, removed := len(arr), 0
	for arrLen > len(arr)/2 {
		arrLen -= xs[0][1]
		xs = xs[1:]
		removed++
	}

	return removed
}
