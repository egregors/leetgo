/*
	https://leetcode.com/problems/h-index/description/

	Given an array of integers citations where citations[i] is the number of
		citations a researcher received for their
	ith paper, return the researcher's h-index.

	According to the definition of h-index on Wikipedia: The h-index is defined as
		the maximum value of h such that
	the given researcher has published at least h papers that have each been cited
		at least h times.
*/

package solutions

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	i := 0
	for i < len(citations) && citations[len(citations)-i-1] > i {
		i++
	}

	return i
}
