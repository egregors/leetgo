/*
	https://leetcode.com/problems/count-largest-group/description

	You are given an integer n.

	Each number from 1 to n is grouped according to the sum of its digits.

	Return the number of groups that have the largest size.
*/

package solutions

func countLargestGroup(n int) int {
	m := make(map[int]int)
	for x := range n {
		m[sum1399(x+1)]++
	}

	gs := 0
	for _, v := range m {
		gs = max(gs, v)
	}

	c := 0
	for _, v := range m {
		if v == gs {
			c++
		}
	}

	return c
}

func sum1399(n int) int {
	s := 0
	for n != 0 {
		s += n % 10
		n /= 10
	}

	return s
}
