/*
	https://leetcode.com/problems/richest-customer-wealth/

	You are given an m x n integer grid accounts where accounts[i][j] is the amount
	of money the ith customer has in the jth bank. Return the wealth that the richest customer has.

	A customer's wealth is the amount of money they have in all their bank accounts.
	The richest customer is the customer that has the maximum wealth.
*/

package solutions

func maximumWealth(accounts [][]int) int {
	max := -1 << 31
	for ri := 0; ri < len(accounts); ri++ {
		if s := Maximum(accounts[ri]...); s > max {
			max = s
		}
	}
	return max
}
