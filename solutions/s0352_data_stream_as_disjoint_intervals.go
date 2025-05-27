/*
	https://leetcode.com/problems/data-stream-as-disjoint-intervals/

	Given a data stream input of non-negative integers a1, a2, ..., an, summarize
		the numbers
	seen so far as a list of disjoint intervals.

	Implement the SummaryRanges class:

		SummaryRanges() Initializes the object with an empty stream.
		void addNum(int value) Adds the integer value to the stream.
		int[][] getIntervals() Returns a summary of the integers in the stream
			currently as a
	list of disjoint intervals [starti, endi]. The answer should be sorted by
		starti.
*/

//nolint:revive // it's ok
package solutions

import "sort"

type SummaryRanges struct {
	vals Set[int]
}

// NewSummaryRanges should call Constructor to pass LeetCode test
func NewSummaryRanges() SummaryRanges {
	return SummaryRanges{vals: Set[int]{}}
}

func (sr *SummaryRanges) AddNum(value int) {
	sr.vals.Add(value)
}

func (sr *SummaryRanges) GetIntervals() [][]int {
	if len(sr.vals) == 0 {
		return [][]int{}
	}
	intervals := [][]int{}
	start, end := -1, -1
	vals := sr.vals.ToSlice()
	sort.Ints(vals)
	for _, val := range vals {
		if start < 0 {
			start = val
			end = val
		} else if val == end+1 {
			end = val
		} else {
			intervals = append(intervals, []int{start, end})
			start = val
			end = val
		}
	}
	intervals = append(intervals, []int{start, end})
	return intervals
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */
