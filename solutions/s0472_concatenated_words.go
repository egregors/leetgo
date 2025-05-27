/*
	https://leetcode.com/problems/concatenated-words/

	Given an array of strings words (without duplicates), return all the
		concatenated
	words in the given list of words.

	A concatenated word is defined as a string that is comprised entirely of at
		least
	two shorter words in the given array.
*/

package solutions

func findAllConcatenatedWordsInADict(words []string) []string {
	m := make(map[string]bool)
	for _, word := range words {
		m[word] = true
	}

	var ans []string
	for _, word := range words {
		length := len(word)
		dp := make([]bool, length+1)
		dp[0] = true
		for i := 1; i <= length; i++ {
			var j int
			if i == length {
				j = 1
			}
			for ; !dp[i] && j < i; j++ {
				dp[i] = dp[j] && m[word[j:i]]
			}
		}
		if dp[length] {
			ans = append(ans, word)
		}
	}
	return ans
}
