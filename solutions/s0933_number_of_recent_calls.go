/*
	https://leetcode.com/problems/number-of-recent-calls/description/

	You have a RecentCounter class which counts the number of recent requests
		within a certain time frame.

	Implement the RecentCounter class:

		RecentCounter() Initializes the counter with zero recent requests.
		int ping(int t) Adds a new request at time t, where t represents some time in
			milliseconds, and returns the
	number of requests that has happened in the past 3000 milliseconds (including
		the new request). Specifically,
	return the number of requests that have happened in the inclusive range [t -
		3000, t].

	It is guaranteed that every call to ping uses a strictly larger value of t than
		the previous call.
*/

package solutions

type RecentCounter struct {
	ts []int
}

// NewRecentCounter should call Constructor to pass LeetCode's test cases
func NewRecentCounter() RecentCounter {
	return RecentCounter{ts: make([]int, 0, 16)}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.ts = append(rc.ts, t)
	for rc.ts[0] < t-3000 {
		rc.ts = rc.ts[1:]
	}

	return len(rc.ts)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
