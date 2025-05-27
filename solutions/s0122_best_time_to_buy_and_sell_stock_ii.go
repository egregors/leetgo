/*
	https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/

	You are given an integer array prices where prices[i] is the price of a given
		stock on the ith day.

	On each day, you may decide to buy and/or sell the stock. You can only hold at
		most one share of the
	stock at any time. However, you can buy it then immediately sell it on the same
		day.

	Find and return the maximum profit you can achieve.
*/

package solutions

func maxProfit0122(prices []int) int {
	s := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			s += prices[i] - prices[i-1]
		}
	}

	return s
}
