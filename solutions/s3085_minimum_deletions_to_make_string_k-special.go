/*
	https://leetcode.com/problems/minimum-deletions-to-make-string-k-special/

	You are given a string word and an integer k.

	We consider word to be k-special if |freq(word[i]) - freq(word[j])| <= k for
		all indices i and j in the string.

	Here, freq(x) denotes the

	of the character x in word, and |y| denotes the absolute value of y.

	Return the minimum number of characters you need to delete to make word
		k-special.
*/

package solutions

func minimumDeletions(word string, k int) int {
	f := make(map[rune]int)
	for _, r := range word {
		f[r]++
	}
	res := len(word)
	for _, v := range f {
		del := 0
		for _, v2 := range f {
			if v > v2 {
				del += v2
			} else if v2 > v+k {
				del += v2 - (v + k)
			}
		}
		if del < res {
			res = del
		}
	}
	return res
}
