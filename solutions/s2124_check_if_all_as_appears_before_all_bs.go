/*
	https://leetcode.com/problems/check-if-all-as-appears-before-all-bs/

	Given a string s consisting of only the characters 'a' and 'b', return true if
		every 'a'
	appears before every 'b' in the string. Otherwise, return false.
*/

package solutions

func checkString(s string) bool {
	flag := false
	curr := rune(s[0])

	for _, ch := range s[1:] {

		if curr == 'b' && ch != curr {
			return false
		}

		if ch != curr {
			if flag {
				return false
			}
			flag, curr = true, ch
		}
	}

	return true
}
