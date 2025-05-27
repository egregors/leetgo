/*
	You are implementing a program to use as your calendar. We can add a new event
		if adding the
	event will not cause a double booking.

	A double booking happens when two events have some non-empty intersection
		(i.e., some moment is
	common to both events.).

	The event can be represented as a pair of integers start and end that
		represents a booking on the
	half-open interval [start, end), the range of real numbers x such that start <=
		x < end.

	Implement the MyCalendar class:

		MyCalendar() Initializes the calendar object.
		boolean book(int start, int end) Returns true if the event can be added to the
			calendar successfully
	without causing a double booking. Otherwise, return false and do not add the
		event to the calendar.
*/

//nolint:revive // it's ok
package solutions

type MyCalendar struct {
	intervals [][]int
}

// NewMyCalendar should call Constructor to pass LeetCode tests
func NewMyCalendar() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start, end int) bool {
	// todo: this is a poor solution.
	//  it could be a BST
	for _, v := range c.intervals {
		if v[0] < end && start < v[1] {
			return false
		}
	}
	c.intervals = append(c.intervals, []int{start, end})
	return true
}
