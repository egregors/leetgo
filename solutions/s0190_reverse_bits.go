/*
	https://leetcode.com/problems/reverse-bits/

	Reverse bits of a given 32 bits unsigned integer.
*/

package solutions

func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	power := uint32(31)
	for num != 0 {
		ret += (num & 1) << power
		num >>= 1
		power--
	}
	return ret
}
