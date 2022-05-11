/*
	https://leetcode.com/problems/count-sorted-vowel-strings/

	Given an integer n, return the number of strings of length n that consist only of
	vowels (a, e, i, o, u) and are lexicographically sorted.

	A string s is lexicographically sorted if for all valid i, s[i] is the same as or comes
	before s[i+1] in the alphabet.
*/

package solutions

func backtrack1641(curr, vowels string, n int, acc *[]string) {
	if len(curr) == n {
		*acc = append(*acc, curr)
		return
	}

	for i := 0; i < len(vowels); i++ {
		if curr == "" || i == 0 {
			curr += string(vowels[i])
		} else {
			if vowels[i] >= vowels[i-1] {
				curr += string(vowels[i])
			} else {
				continue
			}
		}
		backtrack1641(curr, vowels[i:], n, acc)
		curr = curr[:len(curr)-1]
	}
}

func countVowelStrings(n int) int {
	var res []string
	backtrack1641("", "aeiou", n, &res)
	return len(res)
}
