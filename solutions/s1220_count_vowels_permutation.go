/*
	https://leetcode.com/problems/count-vowels-permutation/

		Given an integer n, your task is to count how many strings of length n can be formed under
	the following rules:

		Each character is a lower case vowel ('a', 'e', 'i', 'o', 'u')
		Each vowel 'a' may only be followed by an 'e'.
		Each vowel 'e' may only be followed by an 'a' or an 'i'.
		Each vowel 'i' may not be followed by another 'i'.
		Each vowel 'o' may only be followed by an 'i' or a 'u'.
		Each vowel 'u' may only be followed by an 'a'.

	Since the answer may be too large, return it modulo 10^9 + 7.
*/

package solutions

func countVowelPermutation(n int) int {
	var aCount, eCount, iCount, oCount, uCount int
	aCount, eCount, iCount, oCount, uCount = 1, 1, 1, 1, 1
	MOD := 1000000007

	for i := 1; i < n; i++ {
		aCountNew := (eCount + iCount + uCount) % MOD
		eCountNew := (aCount + iCount) % MOD
		iCountNew := (eCount + oCount) % MOD
		oCountNew := (iCount) % MOD
		uCountNew := (iCount + oCount) % MOD
		aCount = aCountNew
		eCount = eCountNew
		iCount = iCountNew
		oCount = oCountNew
		uCount = uCountNew
	}
	result := (aCount + eCount + iCount + oCount + uCount) % MOD
	return result
}
