/*
	https://leetcode.com/problems/number-of-provinces/

	There are n cities. Some of them are connected, while some are not. If city a is connected
	directly with city b, and city b is connected directly with city c, then city a is connected
	indirectly with city c.

	A province is a group of directly or indirectly connected cities and no other cities outside
	of the group.

	You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and
	the jth city are directly connected, and isConnected[i][j] = 0 otherwise.

	Return the total number of provinces.
*/

package solutions

func dfs547(m [][]int, visited []int, i int) {
	for j := 0; j < len(m); j++ {
		if m[i][j] == 1 && visited[j] == 0 {
			visited[j] = 1
			dfs547(m, visited, j)
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	count := 0
	visited := make([]int, len(isConnected))
	for i := 0; i < len(isConnected); i++ {
		if visited[i] == 0 {
			dfs547(isConnected, visited, i)
			count++
		}
	}
	return count
}
