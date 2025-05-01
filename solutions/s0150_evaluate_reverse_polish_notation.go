/*
	https://leetcode.com/problems/evaluate-reverse-polish-notation/description/

	You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.

	Evaluate the expression. Return an integer that represents the value of the expression.

	Note that:

		The valid operators are '+', '-', '*', and '/'.
		Each operand may be an integer or another expression.
		The division between two integers always truncates toward zero.
		There will not be any division by zero.
		The input represents a valid arithmetic expression in a reverse polish notation.
		The answer and all the intermediate calculations can be represented in a 32-bit integer.

*/

package solutions

import "strconv"

func evalRPN(tokens []string) int {
	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	s := []int{}
	for _, t := range tokens {
		if op, ok := ops[t]; ok {
			a := s[len(s)-2]
			b := s[len(s)-1]
			s = s[:len(s)-2]
			s = append(s, op(a, b))
		} else {
			n, _ := strconv.Atoi(t)
			s = append(s, n)
		}
	}

	return s[0]
}
