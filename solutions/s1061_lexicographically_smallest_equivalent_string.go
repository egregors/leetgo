/*
	https://leetcode.com/problems/lexicographically-smallest-equivalent-string/

	You are given two strings of the same length s1 and s2 and a string baseStr.

	We say s1[i] and s2[i] are equivalent characters.

		For example, if s1 = "abc" and s2 = "cde",
	then we have 'a' == 'c', 'b' == 'd', and 'c' == 'e'.

	Equivalent characters follow the usual rules of any equivalence relation:

		Reflexivity: 'a' == 'a'.
		Symmetry: 'a' == 'b' implies 'b' == 'a'.
		Transitivity: 'a' == 'b' and 'b' == 'c' implies 'a' == 'c'.

	For example, given the equivalency information from s1 = "abc"
	and s2 = "cde", "acd" and "aab" are equivalent strings of
	baseStr = "eed", and "aab" is the lexicographically smallest
	equivalent string of baseStr.

	Return the lexicographically smallest equivalent string of baseStr
	by using the equivalency information from s1 and s2.
*/

package solutions

func smallestEquivalentString(s1, s2, baseStr string) string {
	repr := make([]int, 26)
	var find func(x int) int
	find = func(x int) int {
		if repr[x] == x {
			return x
		}
		repr[x] = find(repr[x])
		return repr[x]
	}

	union := func(x, y int) {
		x = find(x)
		y = find(y)

		if x == y {
			return
		}

		if x < y {
			repr[y] = x
		} else {
			repr[x] = y
		}
	}

	for i := 0; i < 26; i++ {
		repr[i] = i
	}

	for i := 0; i < len(s1); i++ {
		union(int(s1[i]-'a'), int(s2[i]-'a'))
	}
	var ans string
	for i := 0; i < len(baseStr); i++ {
		ans += string(uint8(find(int(baseStr[i]-'a')) + 'a'))
	}
	return ans
}
