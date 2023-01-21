/*
	https://leetcode.com/problems/restore-ip-addresses/description/

	A valid IP address consists of exactly four integers separated by single dots. Each integer is
	between 0 and 255 (inclusive) and cannot have leading zeros.

		For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011.255.245",
	"192.168.1.312" and "192.168@1.1" are invalid IP addresses.

	Given a string s containing only digits, return all possible valid IP addresses that can be
	formed by inserting dots into s. You are not allowed to reorder or remove any digits in s.
	You may return the valid IP addresses in any order.
*/

package solutions

import (
	"strconv"
	"strings"
)

//nolint:revive,stylecheck // it's leetcode signature
func restoreIpAddresses(s string) []string {
	var bt func(string, []string, *[]string)
	bt = func(s string, temp []string, result *[]string) {
		if len(temp) == 4 && len(s) == 0 {
			*result = append(*result, strings.Join(temp, "."))
			return
		}

		if len(temp) == 4 || len(s) > 3*(4-len(temp)) || len(s) < (4-len(temp)) {
			return
		}

		for i := 1; i <= 3; i++ {
			if i > len(s) {
				continue
			}
			num, _ := strconv.Atoi(s[:i])
			if s[:i] != strconv.Itoa(num) || num > 255 {
				continue
			}

			temp = append(temp, s[:i])
			bt(s[i:], temp, result)
			temp = temp[:len(temp)-1]
		}
	}
	var result []string
	bt(s, []string{}, &result)
	return result
}
