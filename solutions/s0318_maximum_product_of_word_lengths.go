/*
	https://leetcode.com/problems/maximum-product-of-word-lengths/

	Given a string array words, return the maximum value of length(word[i]) * length(word[j])
	where the two words do not share common letters. If no such two words exist, return 0.
*/

package solutions

func maxProduct(words []string) int {
	masks := make([]int, len(words))

	for i := 0; i < len(words); i++ {
		bm := 0
		for _, ch := range words[i] {
			bm |= 1 << (ch - 'a')
		}
		masks[i] = bm
	}

	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if masks[i]&masks[j] == 0 {
				max = Maximum(max, len(words[i])*len(words[j]))
			}
		}
	}

	return max
}
