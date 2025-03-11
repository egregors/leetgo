/*
	https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/description/

	Given a string s consisting only of characters a, b and c.

	Return the number of substrings containing at least one occurrence of all these characters a, b and c.
*/

package solutions

func numberOfSubstrings(s string) int {
	pos := []int{-1, -1, -1}
	res := 0
	for i, ch := range s {
		pos[ch-'a'] = i
		res += 1 + min(pos[0], min(pos[1], pos[2]))
	}

	return res
}
