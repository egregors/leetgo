/*
	https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

	You are given an array prices where prices[i] is the price of a given stock on the ith day.

	You want to maximize your profit by choosing a single day to buy one stock and choosing a
	different day in the future to sell that stock.

	Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/

package solutions

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

func maxProfit(prices []int) int {
	buyPrice, maxProf := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > buyPrice {
			maxProf = max(maxProf, prices[i]-buyPrice)
		} else {
			buyPrice = prices[i]
		}
	}
	return maxProf
}
