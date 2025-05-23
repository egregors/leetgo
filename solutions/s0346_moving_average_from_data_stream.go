/*
	https://leetcode.com/problems/moving-average-from-data-stream

	Given a stream of integers and a window size, calculate the moving average of all integers in the sliding window.

	Implement the MovingAverage class:

		MovingAverage(int size) Initializes the object with the size of the window size.
		double next(int val) Returns the moving average of the last size values of the stream.
*/

package solutions

type MovingAverage struct {
	buf []int
	cur int
	cnt int
}

func fsum(xs []int) float64 {
	s := 0
	for _, x := range xs {
		s += x
	}
	return float64(s)
}

// NewMovingAverage must call Constructor to pass LeetCode tests
func NewMovingAverage(size int) MovingAverage {
	return MovingAverage{buf: make([]int, size)}
}

func (a *MovingAverage) Next(val int) float64 {
	a.buf[a.cur] = val
	a.cur = (a.cur + 1) % len(a.buf)
	if a.cnt < len(a.buf) {
		a.cnt++
	}
	return fsum(a.buf) / float64(a.cnt)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
