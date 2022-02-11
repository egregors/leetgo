/*
	https://leetcode.com/problems/generate-parentheses/

	Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
*/

package solutions

func backtrack22(curr []rune, l, r, n int, acc *[][]rune) {
	if len(curr) == n*2 {
		cp := make([]rune, len(curr))
		copy(cp, curr)
		*acc = append(*acc, cp)
		return
	}
	if l < n {
		curr = append(curr, '(')
		backtrack22(curr, l+1, r, n, acc)
		curr = curr[:len(curr)-1]
	}
	if r < l {
		curr = append(curr, ')')
		backtrack22(curr, l, r+1, n, acc)
	}
}

func generateParenthesis(n int) []string {
	var acc [][]rune
	backtrack22([]rune{}, 0, 0, n, &acc)
	res := make([]string, len(acc))
	for i, s := range acc {
		res[i] = string(s)
	}
	return res
}
