/*
	https://leetcode.com/problems/candy/

	There are n children standing in a line. Each child is assigned a rating
	value given in the integer array ratings.

	You are giving candies to these children subjected to the following requirements:

		Each child must have at least one candy.
		Children with a higher rating get more candies than their neighbors.

	Return the minimum number of candies you need to have to distribute the candies
	to the children.
*/

package solutions

func candy(ratings []int) int {
	n := len(ratings)
	l2r, r2l := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		l2r[i], r2l[i] = 1, 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			l2r[i] = l2r[i-1] + 1
		}
	}

	for j := n - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] {
			r2l[j] = r2l[j+1] + 1
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += Maximum(l2r[i], r2l[i])
	}

	return sum
}
