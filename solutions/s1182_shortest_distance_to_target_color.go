/*
	https://leetcode.com/problems/shortest-distance-to-target-color/

	You are given an array colors, in which there are three colors: 1, 2 and 3.

	You are also given some queries. Each query consists of two integers i and c,
	return the shortest distance between the given index i and the target color c.
	If there is no solution return -1.
*/

package solutions

func shortestDistanceColor(colors []int, queries [][]int) []int {
	m := make(map[int][]int)
	for i, c := range colors {
		m[c] = append(m[c], i)
	}

	getIdx := func(i, target int) int {

		ids, ok := m[target]

		if !ok {
			return -1
		}

		var mid int
		lo, hi := 0, len(ids)-1
		for lo <= hi {
			mid = (lo + hi) / 2
			if i == ids[mid] {
				return 0
			}
			if i < ids[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}

		xs := []int{Abs(ids[mid] - i)}
		if mid > 0 {
			xs = append(xs, Abs(ids[mid-1]-i))
		}
		if mid < len(ids)-1 {
			xs = append(xs, Abs(ids[mid+1]-i))
		}

		return Minimum(xs...)
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = getIdx(q[0], q[1])
	}

	return ans
}
