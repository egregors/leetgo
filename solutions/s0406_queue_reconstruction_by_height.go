/*
	https://leetcode.com/problems/queue-reconstruction-by-height/

	You are given an array of people, people, which are the attributes of some
		people in a queue
	(not necessarily in order). Each people[i] = [hi, ki] represents the ith person
		of height hi with
	exactly ki other people in front who have a height greater than or equal to hi.

	Reconstruct and return the queue that is represented by the input array people.
		The returned queue should
	be formatted as an array queue, where queue[j] = [hj, kj] is the attributes of
		the jth person in the queue
	(queue[0] is the person at the front of the queue).
*/

package solutions

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})

	res := make([][]int, len(people))
	empties := make([]int, len(people))

	for i := 0; i < len(people); i++ {
		empties[i] = i
	}

	for _, p := range people {
		res[empties[p[1]]] = p
		empties = append(empties[:p[1]], empties[p[1]+1:]...)
	}

	return res
}
