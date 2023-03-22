/*
	https://leetcode.com/problems/gas-station/

	There are n gas stations along a circular route, where the amount of gas at the ith
	station is gas[i].

	You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the
	ith station to its next (i + 1)th station. You begin the journey with an empty tank at
	one of the gas stations.

	Given two integer arrays gas and cost, return the starting gas station's index if you
	can travel around the circuit once in the clockwise direction, otherwise return -1.
	If there exists a solution, it is guaranteed to be unique
*/

package solutions

func canCompleteCircuit(gas, cost []int) int {
	n := len(gas)
	total, curr := 0, 0
	start := 0
	for i := 0; i < n; i++ {
		total += gas[i] - cost[i]
		curr += gas[i] - cost[i]
		if curr < 0 {
			start = i + 1
			curr = 0
		}
	}
	if total >= 0 {
		return start
	}
	return -1
}
