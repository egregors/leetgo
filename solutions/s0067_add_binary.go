/*
	https://leetcode.com/problems/add-binary

	Given two binary strings a and b, return their sum as a binary string.
*/

package solutions

func addBinary(a string, b string) string {
	var (
		result string
		carry  byte
	)
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		var sum byte
		if i >= 0 {
			sum += a[i] - '0'
		}
		if j >= 0 {
			sum += b[j] - '0'
		}
		sum += carry
		result = string(sum%2+'0') + result
		carry = sum / 2
	}
	return result
}
