/*
	https://leetcode.com/problems/letter-combinations-of-a-phone-number/

	Given a string containing digits from 2-9 inclusive, return all possible letter
		combinations that
	the number could represent. Return the answer in any order.

	A mapping of digit to letters (just like on the telephone buttons) is given
		below.
	Note that 1 does not map to any letters.
*/

package solutions

func backtrack17(idx int, curr []rune, goalLen int, acc *[][]rune, letters [][]rune) {
	if len(curr) == goalLen {
		cp := make([]rune, goalLen)
		copy(cp, curr)
		*acc = append(*acc, cp)
		return
	}

	ls := letters[idx]
	for _, l := range ls {
		curr = append(curr, l)
		backtrack17(idx+1, curr, goalLen, acc, letters)
		curr = curr[:len(curr)-1]
	}
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	ns := map[rune][]rune{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	var letters [][]rune //nolint:prealloc // it's ok
	for _, ch := range digits {
		letters = append(letters, ns[ch])
	}

	var acc [][]rune
	backtrack17(0, []rune{}, len(digits), &acc, letters)

	var res []string //nolint:prealloc // it's ok
	for _, r := range acc {
		res = append(res, string(r))
	}

	return res
}
