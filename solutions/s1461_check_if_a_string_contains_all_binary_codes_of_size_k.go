/*
	https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/

	Given a binary string s and an integer k, return true if every binary code of length k
	is a substring of s. Otherwise, return false.
*/

package solutions

func hasAllCodes(s string, k int) bool {
	need := 1 << k
	got := make([]bool, need)
	allOne := need - 1
	hashVal := 0

	for i := 0; i < len(s); i++ {
		hashVal = ((hashVal << 1) & allOne) | (int(s[i]) - '0')
		if i >= k-1 && !got[hashVal] {
			got[hashVal] = true
			need--
			if need == 0 {
				return true
			}
		}
	}
	return false
}
