/*
	https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/

	You are given a positive integer n representing n cities numbered from 1 to n. You are also given a 2D array roads
	where roads[i] = [ai, bi, distancei] indicates that there is a bidirectional road between cities ai and bi with a
	distance equal to distancei. The cities graph is not necessarily connected.

	The score of a path between two cities is defined as the minimum distance of a road in this path.

	Return the minimum possible score of a path between cities 1 and n.

	Note:

		A path is a sequence of roads between two cities.
		It is allowed for a path to contain the same road multiple times, and you can visit cities 1 and n multiple
	times along the path.
		The test cases are generated such that there is at least one path between 1 and n.
*/

package solutions

import "math"

type pair struct {
	first  int
	second int
}

func minScore(n int, roads [][]int) int {
	adjList := make([][]pair, n+1)
	for _, edge := range roads {
		adjList[edge[0]] = append(adjList[edge[0]], pair{edge[1], edge[2]})
		adjList[edge[1]] = append(adjList[edge[1]], pair{edge[0], edge[2]})
	}
	visited := make([]bool, n+1)
	Queue := []int{1}
	var ans = math.MaxInt
	for len(Queue) > 0 {
		for i := 0; i < len(Queue); i++ {
			var curr = Queue[0]
			Queue = Queue[1:]
			for _, edge := range adjList[curr] {
				if edge.second < ans {
					ans = edge.second
				}
				if !visited[edge.first] {
					Queue = append(Queue, edge.first)
					visited[edge.first] = true
				}
			}
		}
	}
	return ans
}
