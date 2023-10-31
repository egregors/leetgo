/*
	https://leetcode.com/problems/find-the-original-array-of-prefix-xor/description/

	You are given an integer array pref of size n. Find and return the array
	arr of size n that satisfies:

		pref[i] = arr[0] ^ arr[1] ^ ... ^ arr[i].

	Note that ^ denotes the bitwise-xor operation.

	It can be proven that the answer is unique.
*/

package solutions

func findArray(pref []int) []int {
	for i := len(pref) - 1; i > 0; i-- {
		pref[i] ^= pref[i-1]
	}

	return pref
}
