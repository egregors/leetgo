/*
	https://leetcode.com/problems/reverse-vowels-of-a-string/

	Given a string s, reverse only all the vowels in the string and return it.

	The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both
	lower and upper cases, more than once.
*/

package solutions

import "unsafe"

func reverseVowels(s string) string {
	m := map[byte]bool{'a': true, 'e': true, 'i': true,
		'o': true, 'u': true, 'A': true,
		'E': true, 'I': true, 'O': true,
		'U': true}
	bs := []byte(s)
	//nolint:gosec // it's ok to use unsafe here
	str := *(*string)(unsafe.Pointer(&bs))

	lo, hi := 0, len(s)-1
	for lo < hi {
		for lo < len(bs) && !m[bs[lo]] {
			lo++
		}
		for hi >= 0 && !m[bs[hi]] {
			hi--
		}
		if lo > hi {
			return str
		}
		bs[lo], bs[hi] = bs[hi], bs[lo]
		lo++
		hi--
	}
	return str
}
