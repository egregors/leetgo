/*
	https://leetcode.com/problems/minimum-time-to-repair-cars/description/

	You are given an integer array ranks representing the ranks of some mechanics.
		ranksi is the rank of the ith mechanic.
	A mechanic with a rank r can repair n cars in r * n2 minutes.

	You are also given an integer cars representing the total number of cars
		waiting in the garage to be repaired.

	Return the minimum time taken to repair all the cars.

	Note: All the mechanics can repair the cars simultaneously.
*/

package solutions

import "math"

func repairCars(ranks []int, cars int) int64 {
	l, r := 1, int(math.Pow(10, 14))
	for l < r {
		mid := (l + r) >> 1
		if check2594(mid, ranks, cars) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return int64(l)
}

func check2594(mid int, ranks []int, cars int) bool {
	fixed := 0
	for _, r := range ranks {
		fixed += int(math.Sqrt(float64(mid / r)))
		if fixed >= cars {
			return true
		}
	}
	return false
}
