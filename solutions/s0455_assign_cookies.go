/*
https://leetcode.com/problems/assign-cookies/description/

Assume you are an awesome parent and want to give your children some cookies.

	But, you should give each child at most one cookie.

Each child i has a greed factor g[i], which is the minimum size of a cookie that

	the child

	will be content with; and each cookie j has a size s[j]. If s[j] >= g[i], we
		can assign
	the cookie j to the child i, and the child i will be content. Your goal is to
		maximize the
	number of your content children and output the maximum number.
*/
package solutions

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	c := 0
	sort.Ints(g)
	sort.Ints(s)
	for i := len(s) - 1; i >= 0; i-- {
		for j := len(g) - 1; j >= 0; j-- {
			if g[j] != -1 && g[j] <= s[i] {
				c++
				g[j] = -1
				break
			}
		}
	}

	return c
}
