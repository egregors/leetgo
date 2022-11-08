/*
	https://leetcode.com/problems/maximum-69-number/

	You are given a positive integer num consisting only of digits 6 and 9.

	Return the maximum number you can get by changing at most one digit
	(6 becomes 9, and 9 becomes 6).
*/

package solutions

func maximum69Number(num int) int {
	var is []int
	for num/10 > 0 {
		is = append(is, num%10)
		num /= 10
	}
	is = append(is, num%10)
	for i := 0; i < len(is)/2; i++ {
		is[i], is[len(is)-i-1] = is[len(is)-i-1], is[i]
	}

	res := 0
	done := false
	for _, i := range is {
		if !done && i == 6 {
			res *= 10
			res += 9
			done = true
		} else {
			res *= 10
			res += i
		}
	}

	return res
}
