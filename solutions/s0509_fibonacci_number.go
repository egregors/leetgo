/*
	https://leetcode.com/problems/fibonacci-number/

	The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci
	sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1.
	That is,

	F(0) = 0, F(1) = 1
	F(n) = F(n - 1) + F(n - 2), for n > 1.

	Given n, calculate F(n).
*/

package solutions

func dp509(n int, mem map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}
	if _, ok := mem[n]; !ok {
		mem[n] = dp509(n-1, mem) + dp509(n-2, mem)
	}
	return mem[n]
}

func fib(n int) int {
	m := make(map[int]int)
	return dp509(n, m)
}
