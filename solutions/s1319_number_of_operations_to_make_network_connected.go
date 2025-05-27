/*
	https://leetcode.com/problems/number-of-operations-to-make-network-connected

	There are n computers numbered from 0 to n - 1 connected by ethernet cables
		connections forming a
	network where connections[i] = [ai, bi] represents a connection between
		computers ai and bi. Any computer can
	reach any other computer directly or indirectly through the network.

	You are given an initial computer network connections. You can extract certain
		cables between two
	directly connected computers, and place them between any pair of disconnected
		computers to make them directly
	connected.

	Return the minimum number of times you need to do this in order to make all the
		computers connected.
	If it is not possible, return -1.
*/

package solutions

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	var parent []int
	for i := 0; i < n; i++ {
		parent = append(parent, i)
	}

	ans := n - 1

	var find func(a int, parent []int) int
	find = func(a int, parent []int) int {
		if parent[a] == a {
			return a
		}
		res := find(parent[a], parent)
		return res
	}

	union := func(a int, b int, parent []int) {
		parentA := find(a, parent)
		parentB := find(b, parent)

		if parentA != parentB {
			parent[parentA] = parentB
		}
	}

	for _, connection := range connections {
		a := connection[0]
		b := connection[1]
		pa := find(a, parent)
		pb := find(b, parent)

		if pa != pb {
			union(a, b, parent)
			ans--
		}
	}
	return ans
}
