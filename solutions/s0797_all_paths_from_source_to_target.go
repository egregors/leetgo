/*
	https://leetcode.com/problems/all-paths-from-source-to-target/

	Given a directed acyclic graph (DAG) of n nodes labeled from 0 to n - 1, find all possible paths
	from node 0 to node n - 1 and return them in any order.

	The graph is given as follows: graph[i] is a list of all nodes you can visit from
	node i (i.e., there is a directed edge from node i to node graph[i][j]).
*/

package solutions

func backtrack797(curr, target int, graph [][]int, path *[]int, acc *[][]int) {
	if curr == target {
		buf := make([]int, len(*path))
		copy(buf, *path)
		*acc = append(*acc, buf)
		return
	}

	for _, next := range graph[curr] {
		*path = append(*path, next)
		backtrack797(next, target, graph, path, acc)
		*path = (*path)[:len(*path)-1]
	}
}

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	path := []int{0}
	backtrack797(0, len(graph)-1, graph, &path, &res)
	return res
}
