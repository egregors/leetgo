/*
	https://leetcode.com/problems/network-delay-time/

	You are given a network of n nodes, labeled from 1 to n. You are also given times, a list of travel
	times as directed edges times[i] = (ui, vi, wi), where ui is the source node, vi is the target node,
	and wi is the time it takes for a signal to travel from source to target.

	We will send a signal from a given node k. Return the time it takes for all the n nodes to receive the
	signal. If it is impossible for all the n nodes to receive the signal, return -1.
*/

package solutions

import "math"

func networkDelayTime(times [][]int, n, k int) int {
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0

	for i := 0; i < n; i++ {
		for _, time := range times {
			u := time[0]
			v := time[1]
			w := time[2]
			if dist[u] != math.MaxInt32 && dist[v] > dist[u]+w {
				dist[v] = dist[u] + w
			}
		}
	}

	maxWait := 0
	for x := 1; x <= n; x++ {
		maxWait = int(math.Max(float64(maxWait), float64(dist[x])))
	}

	if maxWait != math.MaxInt32 {
		return maxWait
	}
	return -1
}
