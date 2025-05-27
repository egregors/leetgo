/*
	https://leetcode.com/problems/koko-eating-bananas

	Koko loves to eat bananas. There are n piles of bananas, the ith pile has
		piles[i]
	bananas. The guards have gone and will come back in h hours.

	Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses
	some pile of bananas and eats k bananas from that pile. If the pile has less
	than k bananas, she eats all of them instead and will not eat any more bananas
	during this hour.

	Koko likes to eat slowly but still wants to finish eating all the bananas
		before
	the guards return.

	Return the minimum integer k such that she can eat all the bananas within h
		hours.
*/

package solutions

import "math"

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, 1
	for _, p := range piles {
		hi = Maximum(hi, p)
	}

	for lo < hi {
		mid := (lo + hi) / 2
		var spent int

		for _, p := range piles {
			spent += int(math.Ceil(float64(p) / float64(mid)))
		}

		if spent <= h {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return hi
}
