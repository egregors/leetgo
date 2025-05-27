/*
	https://leetcode.com/problems/basic-calculator-ii/description/

	Given a string s which represents an expression, evaluate this expression and
		return its value.

	The integer division should truncate toward zero.

	You may assume that the given expression is always valid. All intermediate
		results will be in the
	range of [-231, 231 - 1].

	Note: You are not allowed to use any built-in function which evaluates strings
		as mathematical
	expressions, such as eval().
*/

package solutions

import "unicode"

func calculate(s string) int {
	stack := &[]int{}
	op := '+'
	curNum := 0

	for i, r := range s {
		if unicode.IsDigit(r) {
			curNum = (curNum * 10) + (int(r) - '0')
		}
		if !unicode.IsDigit(r) && r != ' ' || i == len(s)-1 {
			switch op {
			case '+':
				push(stack, curNum)
			case '-':
				push(stack, -curNum)
			case '*':
				push(stack, pop(stack)*curNum)
			case '/':
				push(stack, pop(stack)/curNum)
			}
			op = r
			curNum = 0
		}
	}

	sum := 0
	for len(*stack) > 0 {
		sum += pop(stack)
	}

	return sum
}

func pop(xs *[]int) int {
	x := (*xs)[len(*xs)-1]
	*xs = (*xs)[:len(*xs)-1]
	return x
}

func push(xs *[]int, x int) {
	*xs = append(*xs, x)
}
