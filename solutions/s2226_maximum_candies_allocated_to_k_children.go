/*
	https://leetcode.com/problems/maximum-candies-allocated-to-k-children/description/

	You are given a 0-indexed integer array candies. Each element in the array
		denotes a pile of candies
	of size candies[i]. You can divide each pile into any number of sub piles, but
		you cannot merge two piles together.

	You are also given an integer k. You should allocate piles of candies to k
		children such that each child
	gets the same number of candies. Each child can be allocated candies from only
		one pile of candies and some piles of candies may go unused.

	Return the maximum number of candies each child can get.
*/

package solutions

import "sort"

func maximumCandies(candies []int, k int64) int {
	if Sum(candies...)/int(k) == 0 {
		return 0
	}

	sort.Ints(candies)
	l, r := 1, candies[len(candies)-1]

	var mid int
	for l <= r {
		mid = (l + r) >> 1
		if check(mid, k, candies) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l - 1
}

func check(c int, k int64, xs []int) bool {
	n := 0
	for i := 0; i < len(xs); i++ {
		n += xs[i] / c
		if n >= int(k) {
			return true
		}
	}

	return false
}
