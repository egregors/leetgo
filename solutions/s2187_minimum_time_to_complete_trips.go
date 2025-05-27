/*
	https://leetcode.com/problems/minimum-time-to-complete-trips/

	You are given an array time where time[i] denotes the time taken by the ith bus
		to
	complete one trip.

	Each bus can make multiple trips successively; that is, the next trip can start
	immediately after completing the current trip. Also, each bus operates
		independently;
	that is, the trips of one bus do not influence the trips of any other bus.

	You are also given an integer totalTrips, which denotes the number of trips all
		buses
	should make in total. Return the minimum time required for all buses to
		complete at
	least totalTrips trips.
*/

package solutions

func minimumTime(time []int, totalTrips int) int64 {
	isTimeEnough := func(time []int, given, totalTrips int) bool {
		actualT := 0
		for _, t := range time {
			actualT += given / t
		}
		return actualT >= totalTrips
	}

	max := 0
	for _, t := range time {
		max = Maximum(max, t)
	}
	lo, hi := 1, max*totalTrips

	for lo < hi {
		mid := (lo + hi) >> 1
		if isTimeEnough(time, mid, totalTrips) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return int64(lo)
}
