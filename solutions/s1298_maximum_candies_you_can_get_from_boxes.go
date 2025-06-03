/*
	https://leetcode.com/problems/maximum-candies-you-can-get-from-boxes/

	You have n boxes labeled from 0 to n - 1. You are given four arrays: status,
		candies, keys, and containedBoxes where:

		status[i] is 1 if the ith box is open and 0 if the ith box is closed,
		candies[i] is the number of candies in the ith box,
		keys[i] is a list of the labels of the boxes you can open after opening the
			ith box.
		containedBoxes[i] is a list of the boxes you found inside the ith box.

	You are given an integer array initialBoxes that contains the labels of the
		boxes you initially have. You can take all the candies in any open box and you
		can use the keys in it to open new boxes and you also can use the boxes you
		find in it.

	Return the maximum number of candies you can get following the rules above.
*/

package solutions

import (
	"sort"
)

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	cnt := 0
	ks := make(map[int]bool)
	circle := 0

	eq := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		sort.Ints(a)
		sort.Ints(b)
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}

		return true
	}

	q := initialBoxes
	for len(q) > 0 {
		nextQ := []int{}
		for _, box := range q {
			if status[box] == 1 { // box opened
				// get candies
				cnt += candies[box]
				// get keys
				for _, k := range keys[box] {
					ks[k] = true
				}
				// get new boxes
				nextQ = append(nextQ, containedBoxes[box]...)
			} else { // box closed
				if ks[box] {
					status[box] = 1
				}
				nextQ = append(nextQ, box)
			}
		}

		if eq(q, nextQ) {
			circle++
		} else {
			circle = 0
		}
		if circle >= 3 {
			break
		}
		q = nextQ
	}

	return cnt
}
