/*
	https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/description/

	You are given a 0-indexed integer array tasks, where tasks[i] represents the
		difficulty
	level of a task. In each round, you can complete either 2 or 3 tasks of the
		same difficulty
	level.

	Return the minimum rounds required to complete all the tasks, or -1 if it is
		not possible
	to complete all the tasks.
*/

package solutions

func minimumRounds(tasks []int) int {
	m := make(map[int]int)
	for _, t := range tasks {
		m[t]++
	}

	minRounds := 0
	for _, v := range m {
		if v == 1 {
			return -1
		}

		if v%3 == 0 {
			minRounds += v / 3
		} else {
			minRounds += v/3 + 1
		}
	}

	return minRounds
}
