/*
	https://leetcode.com/problems/high-five/

	Given a list of the scores of different students, items, where items[i] = [IDi,
		scorei] represents one score from a student with IDi, calculate each student's
		top five average.

	Return the answer as an array of pairs result, where result[j] = [IDj,
		topFiveAveragej] represents the student with IDj and their top five average.
		Sort result by IDj in increasing order.

	A student's top five average is calculated by taking the sum of their top five
		scores and dividing it by 5 using integer division.
*/

package solutions

import "sort"

func highFive(items [][]int) [][]int {
	m := make(map[int][]int)
	for _, item := range items {
		id, score := item[0], item[1]
		m[id] = append(m[id], score)
	}
	for k := range m {
		sort.Ints(m[k])
	}

	sum := func(xs ...int) int {
		s := 0
		for _, x := range xs {
			s += x
		}
		return s
	}

	res := make([][]int, 0, len(m))
	for k, vs := range m {
		res = append(res, []int{
			k,
			sum(vs[len(vs)-5:]...) / 5,
		})
	}

	sort.Slice(res, func(i, j int) bool { return res[i][0] < res[j][0] })

	return res
}
