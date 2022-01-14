/*
	https://leetcode.com/problems/string-to-integer-atoi/

	Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer
	(similar to C/C++'s atoi function).

	The algorithm for myAtoi(string s) is as follows:

		Read in and ignore any leading whitespace.
		Check if the next character (if not already at the end of the string) is '-' or '+'.
		Read this character in if it is either. This determines if the final result is negative
		or positive respectively. Assume the result is positive if neither is present.
		Read in next the characters until the next non-digit character or the end of the input
		is reached. The rest of the string is ignored.
		Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits
		were read, then the integer is 0. Change the sign as necessary (from step 2).
		If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp
		the integer so that it remains in the range. Specifically, integers less than -231 should
		be clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
		Return the integer as the final result.

	Note:

		Only the space character ' ' is considered a whitespace character.
		Do not ignore any characters other than the leading whitespace or the rest of the string
		after the digits.
*/

package solutions

import (
	"math"
)

const (
	ws    = ' '
	zero  = '0'
	plus  = '+'
	minus = '-'
)

func parseInt(bs *[]byte) int {
	if len(*bs) == 0 {
		return 0
	}

	nums := map[byte]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}

	// get sign
	isNegative := !parseSign(bs)

	// cut tail
	for i := 0; i < len(*bs); i++ {
		if !nums[(*bs)[i]] {
			*bs = (*bs)[:i]
			break
		}
	}

	// trim zeros
	for len(*bs) > 0 && (*bs)[0] == zero {
		*bs = (*bs)[1:]
	}

	// check Int size and clamp
	var un int
	for i := 0; i < len(*bs); i++ {
		un *= 10
		un1 := un + int((*bs)[i]) - '0'

		if un1 > math.MaxInt32 || (-1)*un1 < math.MinInt32 {
			if isNegative {
				return math.MinInt32
			}
			return math.MaxInt32
		}

		un = un1
	}

	if isNegative {
		un *= -1
	}

	return un
}

// parseSign checks if here a sign, if so – remove it.
//      if positive - true
//      if negative - false
func parseSign(bs *[]byte) bool {
	if (*bs)[0] == minus {
		*bs = (*bs)[1:]
		return false
	}
	if (*bs)[0] == plus {
		*bs = (*bs)[1:]
	}
	return true
}

func trimLeft(bs *[]byte) {
	for len(*bs) > 0 && (*bs)[0] == ws {
		*bs = (*bs)[1:]
	}
}

func myAtoi(s string) int {
	if s == "" {
		return 0
	}

	bs := []byte(s)
	trimLeft(&bs)
	res := parseInt(&bs)
	return res
}
