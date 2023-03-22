/*
	https://leetcode.com/problems/word-subsets/

	You are given two string arrays words1 and words2.

	A string b is a subset of string a if every letter in b occurs in a including multiplicity.

		For example, "wrr" is a subset of "warrior" but is not a subset of "world".

	A string a from words1 is universal if for every string b in words2, b is a subset of a.

	Return an array of all the universal strings in words1. You may return the answer in any order.
*/

package solutions

func wordSubsets(words1, words2 []string) []string {
	commonSub := make([]int, 26)
	for _, w := range words2 {
		arr := make([]int, 26)
		for _, ch := range w {
			arr[ch-'a']++
		}
		for i, ch := range arr {
			commonSub[i] = Maximum(commonSub[i], ch)
		}
	}

	isSubStr := func(w string) bool {
		wArr := make([]int, 26)
		for _, ch := range w {
			wArr[ch-'a']++
		}
		for i := 0; i < 26; i++ {
			if commonSub[i] > 0 {
				if wArr[i]-commonSub[i] < 0 {
					return false
				}
			}
		}
		return true
	}

	var res []string
	for _, w := range words1 {
		if isSubStr(w) {
			res = append(res, w)
		}
	}
	return res
}
