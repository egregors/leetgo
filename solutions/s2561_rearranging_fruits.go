/*
	https://leetcode.com/problems/rearranging-fruits

	You have two fruit baskets containing n fruits each. You are given two
		0-indexed integer arrays basket1 and basket2 representing the cost of fruit in
		each basket. You want to make both baskets equal. To do so, you can use the
		following operation as many times as you want:

		Chose two indices i and j, and swap the ith fruit of basket1 with the jth
			fruit of basket2.
		The cost of the swap is min(basket1[i],basket2[j]).

	Two baskets are considered equal if sorting them according to the fruit cost
		makes them exactly the same baskets.

	Return the minimum cost to make both the baskets equal or -1 if impossible.
*/

package solutions

import (
	"math"
	"sort"
)

func minCost2561(basket1 []int, basket2 []int) int64 {
	freq := map[int]int{}
	m := math.MaxInt
	for _, b := range basket1 {
		freq[b]++
		if b < m {
			m = b
		}
	}
	for _, b := range basket2 {
		freq[b]--
		if b < m {
			m = b
		}
	}

	var merge []int
	for k, c := range freq {
		if c%2 != 0 {
			return -1
		}
		for i := 0; i < Abs(c)/2; i++ {
			merge = append(merge, k)
		}
	}

	sort.Ints(merge)
	var res int64
	for i := 0; i < len(merge)/2; i++ {
		if 2*m < merge[i] {
			res += int64(2 * m)
		} else {
			res += int64(merge[i])
		}
	}
	return res
}
