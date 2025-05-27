/*
	https://leetcode.com/problems/ipo/

	Suppose LeetCode will start its IPO soon. In order to sell a good price of its
	shares to Venture Capital, LeetCode would like to work on some projects to
	increase its capital before the IPO. Since it has limited resources, it can
		only
	finish at most k distinct projects before the IPO. Help LeetCode design the
		best
	way to maximize its total capital after finishing at most k distinct projects.

	You are given n projects where the ith project has a pure profit profits[i] and
	a minimum capital of capital[i] is needed to start it.

	Initially, you have w capital. When you finish a project, you will obtain its
	pure profit and the profit will be added to your total capital.

	Pick a list of at most k distinct projects from given projects to maximize your
	final capital, and return the final maximized capital.

	The answer is guaranteed to fit in a 32-bit signed integer.
*/

//nolint:revive // it's ok
package solutions

import (
	"container/heap"
	"sort"
)

type Project struct {
	profit, capital int
}

func findMaximizedCapital(k, w int, profits, capital []int) int {
	n := len(profits)
	projects := make([]Project, n)
	for i := range profits {
		projects[i] = Project{profits[i], capital[i]}
	}
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	q := &IntMaxHeap{}
	heap.Init(q)

	ptr := 0
	for i := 0; i < k; i++ {
		for ptr < n && projects[ptr].capital <= w {
			heap.Push(q, projects[ptr].profit)
			ptr++
		}
		if q.Len() == 0 {
			break
		}
		w += heap.Pop(q).(int)
	}
	return w
}
