/*
	https://leetcode.com/problems/finding-3-digit-even-numbers

	You are given an integer array digits, where each element is a digit. The array may contain duplicates.

	You need to find all the unique integers that follow the given requirements:

		The integer consists of the concatenation of three elements from digits in any arbitrary order.
		The integer does not have leading zeros.
		The integer is even.

	For example, if the given digits were [1, 2, 3], integers 132 and 312 follow the requirements.

	Return a sorted array of the unique integers.
*/

package solutions

func findEvenNumbers(digits []int) []int {
	m := make([]int, 10)
	for _, d := range digits {
		m[d]++
	}

	pred := func(n int, m []int) bool {
		nums := []int{}
		for n != 0 {
			nums = append(nums, n%10)
			n /= 10
		}

		used := make([]int, 10)
		for _, n := range nums {
			used[n]++
		}

		for i := range used {
			if m[i]-used[i] < 0 {
				return false
			}
		}

		return true
	}

	res := make([]int, 0, 1024)
	for i := 100; i < 1000; i += 2 {
		if pred(i, m) {
			res = append(res, i)
		}
	}

	return res
}
