/*
	https://leetcode.com/problems/find-and-replace-pattern/

	Given a list of strings words and a string pattern, return a list of words[i]
		that match pattern.
	You may return the answer in any order.

	A word matches the pattern if there exists a permutation of letters p so that
		after replacing every
	letter x in the pattern with p(x), we get the desired word.

	Recall that a permutation of letters is a bijection from letters to letters:
		every letter maps to
	another letter, and no two letters map to the same letter.
*/

package solutions

func findAndReplacePattern(words []string, pattern string) []string {
	isValid := func(pattern, s string) bool {
		if len(pattern) != len(s) {
			return false
		}

		m := make(map[byte]byte)
		for i := 0; i < len(s); i++ {
			w, p := s[i], pattern[i]
			if v, ok := m[w]; ok {
				if v != p {
					return false
				}
			} else {
				m[w] = p
			}
		}

		seen := make([]bool, 26)
		for _, v := range m {
			if seen[v-'a'] {
				return false
			}
			seen[v-'a'] = true
		}
		return true
	}
	return filter(func(el string) bool { return isValid(pattern, el) }, words)
}

func filter(pred func(string) bool, ws []string) []string {
	var res []string
	for _, w := range ws {
		if pred(w) {
			res = append(res, w)
		}
	}
	return res
}
