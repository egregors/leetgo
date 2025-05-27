/*
	https://leetcode.com/problems/minimum-cost-for-tickets/

		You have planned some train traveling one year in advance. The days of the
			year in which you will travel
	are given as an integer array days. Each day is an integer from 1 to 365.

	Train tickets are sold in three different ways:

		a 1-day pass is sold for costs[0] dollars,
		a 7-day pass is sold for costs[1] dollars, and
		a 30-day pass is sold for costs[2] dollars.

	The passes allow that many days of consecutive travel.

		For example, if we get a 7-day pass on day 2, then we can travel for 7 days:
			2, 3, 4, 5, 6, 7, and 8.

	Return the minimum number of dollars you need to travel every day in the given
		list of days.
*/

package solutions

func mincostTickets(days, costs []int) int {
	n := days[len(days)-1]
	dp := make([]int, n+1)
	M := make(map[int]bool)
	for _, v := range days {
		M[v] = true
	}
	for i := 1; i < len(dp); i++ {
		if !M[i] {
			dp[i] = dp[i-1]
			continue
		}
		one := dp[i-1] + costs[0]
		seven := dp[Maximum(i-7, 0)] + costs[1]
		thirty := dp[Maximum(i-30, 0)] + costs[2]
		dp[i] = Minimum(one, seven, thirty)
	}
	return dp[n]
}
