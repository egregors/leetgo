/*
	https://leetcode.com/problems/online-stock-span/

	Design an algorithm that collects daily price quotes for some stock and returns
		the span of that stock's
	price for the current day.

	The span of the stock's price today is defined as the maximum number of
		consecutive days (starting from today
	and going backward) for which the stock price was less than or equal to today's
		price.

		For example, if the price of a stock over the next 7 days were
			[100,80,60,70,60,75,85], then the stock spans
		would be [1,1,1,2,1,4,6].

	Implement the StockSpanner class:

		StockSpanner() Initializes the object of the class.
		int next(int price) Returns the span of the stock's price given that today's
			price is price.
*/

//nolint:revive // it's ok
package solutions

type StockSpanner struct {
	xs []int
}

// NewStockSpanner should call Constructor() to pass LeetCode tests
func NewStockSpanner() StockSpanner {
	return StockSpanner{xs: []int{}}
}

func (s *StockSpanner) Next(price int) int {
	s.xs = append(s.xs, price)
	i := len(s.xs) - 1
	for i >= 0 && s.xs[i] <= price {
		i--
	}
	return len(s.xs) - 1 - i
}
