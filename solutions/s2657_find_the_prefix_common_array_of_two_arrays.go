/*
	2657. Find the Prefix Common Array of Two Arrays

	You are given two 0-indexed integer permutations A and B of length n.

	A prefix common array of A and B is an array C such that C[i] is equal to the
		count of numbers that are
	present at or before the index i in both A and B.

	Return the prefix common array of A and B.

	A sequence of n integers is called a permutation if it contains all integers
		from 1 to n exactly once.
*/

package solutions

func findThePrefixCommonArray(A []int, B []int) []int {
	xs := make([]int, len(A)+1)
	c := 0
	res := make([]int, 0, len(A))
	for i := 0; i < len(A); i++ {
		xs[A[i]]++
		if xs[A[i]] == 2 {
			c++
		}
		xs[B[i]]++
		if xs[B[i]] == 2 {
			c++
		}
		res = append(res, c)
	}

	return res
}
