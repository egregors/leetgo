/*
	https://leetcode.com/problems/minimum-fuel-cost-to-report-to-the-capital/

	There is a tree (i.e., a connected, undirected graph with no cycles) structure country
	network consisting of n cities numbered from 0 to n - 1 and exactly n - 1 roads.
	The capital city is city 0. You are given a 2D integer array roads where roads[i] =
	[ai, bi] denotes that there exists a bidirectional road connecting cities ai and bi.

	There is a meeting for the representatives of each city. The meeting is in the capital
	city.

	There is a car in each city. You are given an integer seats that indicates the number
	of seats in each car.

	A representative can use the car in their city to travel or change the car and ride
	with another representative. The cost of traveling between two cities is one liter of fuel.

	Return the minimum number of liters of fuel to reach the capital city.
*/

package solutions

import "math"

func minimumFuelCost(roads [][]int, seats int) int64 {
	var ans int64

	var dfs func(int, int, map[int][]int, int) int
	dfs = func(node, parent int, adj map[int][]int, seats int) int {
		rep := 1
		if _, ok := adj[node]; !ok {
			return rep
		}
		for _, child := range adj[node] {
			if child != parent {
				rep += dfs(child, node, adj, seats)
			}
		}
		if node != 0 {
			ans += int64(math.Ceil(float64(rep) / float64(seats)))
		}
		return rep
	}

	adj := make(map[int][]int)
	for _, r := range roads {
		adj[r[0]] = append(adj[r[0]], r[1])
		adj[r[1]] = append(adj[r[1]], r[0])
	}
	dfs(0, -1, adj, seats)
	return ans
}
