/*
	https://leetcode.com/problems/sum-of-subarray-minimums/

	Given an array of integers arr, find the sum of min(b), where b ranges over every (contiguous) subarray of arr.
	Since the answer may be large, return the answer modulo 109 + 7.
*/

//nolint:revive // it's ok
package solutions

type Stack907 []int

func (s Stack907) IsEmpty() bool { return len(s) == 0 }
func (s Stack907) Peek() int     { return s[len(s)-1] }
func (s *Stack907) Push(x int)   { *s = append(*s, x) }
func (s *Stack907) Pop()         { *s = (*s)[:len(*s)-1] }

func sumSubarrayMins(arr []int) int {
	dp := make([]int, len(arr))

	s := Stack907{}

	for i := 0; i < len(arr); i++ {
		for !s.IsEmpty() && arr[s.Peek()] >= arr[i] {
			s.Pop()
		}
		if !s.IsEmpty() {
			prev := s.Peek()
			dp[i] = dp[prev] + (i-prev)*arr[i]
		} else {
			dp[i] = (i + 1) * arr[i]
		}
		s.Push(i)
	}

	res := 0
	for _, n := range dp {
		res += n
		res %= 1000000007
	}

	return res
}
