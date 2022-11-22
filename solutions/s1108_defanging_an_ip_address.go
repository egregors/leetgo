/*
	https://leetcode.com/problems/defanging-an-ip-address/

	Given a valid (IPv4) IP address, return a defanged version of that IP address.

	A defanged IP address replaces every period "." with "[.]".
*/

package solutions

func defangIPaddr(address string) string {
	def := []rune("[.]")
	var res []rune
	for _, r := range address {
		if r == '.' {
			res = append(res, def...)
		} else {
			res = append(res, r)
		}
	}
	return string(res)
}
