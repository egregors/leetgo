/*
	https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/

	Given an array of integers arr and an integer k. Find the least number of
		unique
	integers after removing exactly k elements.
*/

package solutions

import (
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	xs := make(map[int]int)
	keys := make([]int, 0, len(xs))

	for _, n := range arr {
		if _, ok := xs[n]; !ok {
			keys = append(keys, n)
		}
		xs[n]++
	}

	sort.Slice(keys, func(i, j int) bool {
		return xs[keys[i]] < xs[keys[j]]
	})

	for i := 0; k > 0 && i < len(keys); i++ {
		n := keys[i]
		if xs[n]-k < 1 {
			k -= xs[n]
			delete(xs, n)
		}
	}

	return len(xs)
}
