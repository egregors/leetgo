/*
	https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/description/

	A happy string is a string that:

		consists only of letters of the set ['a', 'b', 'c'].
		s[i] != s[i + 1] for all values of i from 1 to s.length - 1 (string is
			1-indexed).

	For example, strings "abc", "ac", "b" and "abcbabcbcb" are all happy strings
		and strings "aa", "baa" and "ababbc" are not happy strings.

	Given two integers n and k, consider a list of all happy strings of length n
		sorted in lexicographical order.

	Return the kth string of this list or return an empty string if there are less
		than k happy strings of length n.
*/

package solutions

import "sort"

func getHappyString(n int, k int) string {
	res := make([]string, 0, k)

	var bt func(curr []rune)
	bt = func(curr []rune) {
		if len(curr) == n {
			res = append(res, string(curr))
			return
		}

		for ch := 'a'; ch <= 'c'; ch++ {
			if len(curr) > 0 && curr[len(curr)-1] == ch {
				continue
			}

			bt(append(curr, ch))
		}
	}

	bt(make([]rune, 0, k))
	sort.Strings(res)
	if k > len(res) {
		return ""
	}

	return res[k-1]
}
