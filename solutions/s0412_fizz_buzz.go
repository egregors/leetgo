/*
	https://leetcode.com/problems/fizz-buzz/

	Given an integer n, return a string array answer (1-indexed) where:

    answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
    answer[i] == "Fizz" if i is divisible by 3.
    answer[i] == "Buzz" if i is divisible by 5.
    answer[i] == i (as a string) if none of the above conditions are true.
*/

package solutions

import "strconv"

func fizzBuzz(n int) []string {
	res := []string{}

	for i := 1; i <= n; i++ {
		cur := ""
		if i%3 == 0 {
			cur += "Fizz"
		}
		if i%5 == 0 {
			cur += "Buzz"
		}
		if cur == "" {
			cur = strconv.Itoa(i)
		}
		res = append(res, cur)
	}

	return res
}
