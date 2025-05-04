/*
	https://leetcode.com/problems/number-of-equivalent-domino-pairs/description/

	Given a list of dominoes, dominoes[i] = [a, b] is equivalent to dominoes[j] = [c, d] if and only if either
	(a == c and b == d), or (a == d and b == c) - that is, one domino can be rotated to be equal to another domino.

	Return the number of pairs (i, j) for which 0 <= i < j < dominoes.length, and dominoes[i] is equivalent to
	dominoes[j].
*/

package solutions

func numEquivDominoPairs(dominoes [][]int) int {
	cnt := [100]int{}
	ans := 0
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		v := d[0]*10 + d[1]
		ans += cnt[v]
		cnt[v]++
	}
	return ans
}
