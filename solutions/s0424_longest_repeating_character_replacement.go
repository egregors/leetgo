/*
	https://leetcode.com/problems/longest-repeating-character-replacement/description

	You are given a string s and an integer k. You can choose any character of the string and change it to any other
	uppercase English character. You can perform this operation at most k times.

	Return the length of the longest substring containing the same letter you can get after performing the above
	operations.
*/

package solutions

func characterReplacement(s string, k int) int {
	start := 0
	fMap := make([]int, 26)
	maxF := 0
	ans := 0

	for end := 0; end < len(s); end++ {
		ch := s[end] - 'A'
		fMap[ch]++
		maxF = max(maxF, fMap[ch])

		if !(end-start+1-maxF <= k) { //nolint:staticcheck // meh
			chToDrop := s[start] - 'A'
			fMap[chToDrop]--
			start++
		}

		ans = end - start + 1
	}

	return ans
}
