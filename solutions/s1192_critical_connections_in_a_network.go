/*
	https://leetcode.com/problems/critical-connections-in-a-network/

	There are n servers numbered from 0 to n - 1 connected by undirected server-to-server
	connections forming a network where connections[i] = [ai, bi] represents a connection
	between servers ai and bi. Any server can reach other servers directly or indirectly
	through the network.

	A critical connection is a connection that, if removed, will make some servers unable
	to reach some other server.

	Return all critical connections in the network in any order.
*/

package solutions

func criticalConnections(n int, connections [][]int) [][]int {
	graph := make([][]int, n)

	for _, c := range connections {
		graph[c[0]] = append(graph[c[0]], c[1])
		graph[c[1]] = append(graph[c[1]], c[0])
	}

	return dfs1192(0, -1, 0, graph, make([]*int, n))
}

func dfs1192(cur, prev, depth int, graph [][]int, visited []*int) (res [][]int) {
	visited[cur] = &depth

	for _, n := range graph[cur] {

		if n == prev {
			continue
		}

		if visited[n] == nil {
			res = append(res, dfs1192(n, cur, depth+1, graph, visited)...)
		}

		if *visited[cur] > *visited[n] {
			visited[cur] = visited[n]
		}

		if *visited[n] > depth {
			res = append(res, []int{cur, n})
		}
	}

	return res
}
