/*
	https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/

	A conveyor belt has packages that must be shipped from one port to another
		within days days.

	The ith package on the conveyor belt has a weight of weights[i]. Each day, we
		load the ship with
	packages on the conveyor belt (in the order given by weights). We may not load
		more weight than the
	maximum weight capacity of the ship.

	Return the least weight capacity of the ship that will result in all the
		packages on the conveyor belt
	being shipped within days days.
*/

package solutions

func shipWithinDays(weights []int, days int) int {
	feasible := func(weights []int, c, days int) bool {
		daysNeeded, currentLoad := 1, 0
		for _, weight := range weights {
			currentLoad += weight
			if currentLoad > c {
				daysNeeded++
				currentLoad = weight
			}
		}

		return daysNeeded <= days
	}

	var totalLoad, maxLoad int
	for _, weight := range weights {
		totalLoad += weight
		maxLoad = Maximum(maxLoad, weight)
	}
	l, r := maxLoad, totalLoad

	for l < r {
		mid := (l + r) / 2
		if feasible(weights, mid, days) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
