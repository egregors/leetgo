/*
	https://leetcode.com/problems/maximize-distance-to-closest-person/description/

	You are given an array representing a row of seats where seats[i] = 1
		represents a person sitting in the ith seat,
	and seats[i] = 0 represents that the ith seat is empty (0-indexed).

	There is at least one empty seat, and at least one person sitting.

	Alex wants to sit in the seat such that the distance between him and the
		closest person to him is maximized.

	Return that maximum distance to the closest person.
*/

package solutions

func maxDistToClosest(seats []int) int {
	n := len(seats)
	var l, r []int //nolint:prealloc //meh
	for range n {
		r = append(r, n)
		l = append(l, n)
	}

	for i := 0; i < n; i++ {
		if seats[i] == 1 {
			l[i] = 0
		} else if i > 0 {
			l[i] = l[i-1] + 1
		}
	}

	for i := n - 1; i >= 0; i-- {
		if seats[i] == 1 {
			r[i] = 0
		} else if i < n-1 {
			r[i] = r[i+1] + 1
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		if seats[i] == 0 {
			res = max(res, min(l[i], r[i]))
		}
	}

	return res
}
