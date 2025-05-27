/*
	A string is good if there are no repeated characters.

	Given a string s, return the number of good substrings of
	length three in s.

	Note that if there are multiple occurrences of the same substring, every
		occurrence
	should be counted.

	A substring is a contiguous sequence of characters in a string.
*/

package solutions

func countGoodSubstrings(s string) (counter int) {
	for i := 0; i < len(s)-2; i++ {
		if s[i] != s[i+1] && s[i+1] != s[i+2] && s[i] != s[i+2] {
			counter++
		}
	}
	return
}
