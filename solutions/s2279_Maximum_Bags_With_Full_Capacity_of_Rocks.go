/*
	https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/description/

	You have n bags numbered from 0 to n - 1. You are given two 0-indexed integer arrays
	capacity and rocks. The ith bag can hold a maximum of capacity[i] rocks and currently
	contains rocks[i] rocks. You are also given an integer additionalRocks, the number of
	additional rocks you can place in any of the bags.

	Return the maximum number of bags that could have full capacity after placing the
	additional rocks in some bags.
*/

//nolint:revive // it's ok
package solutions

import "container/heap"

type Bag struct {
	len, cap int
}

type Bags []Bag

func (bs Bags) Len() int      { return len(bs) }
func (bs Bags) IsEmpty() bool { return bs.Len() == 0 }
func (bs Bags) Less(i, j int) bool {
	return (bs[i].cap - bs[i].len) < (bs[j].cap - bs[j].len)
}

func (bs *Bags) Swap(i, j int) {
	(*bs)[i], (*bs)[j] = (*bs)[j], (*bs)[i]
}

func (bs *Bags) Push(x interface{}) {
	*bs = append(*bs, x.(Bag))
}

func (bs *Bags) Pop() interface{} {
	x := (*bs)[bs.Len()-1]
	*bs = (*bs)[:bs.Len()-1]
	return x
}

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	bags := Bags{}
	for i := range rocks {
		heap.Push(&bags, Bag{len: rocks[i], cap: capacity[i]})
	}

	var res int
	for additionalRocks > 0 && !bags.IsEmpty() {
		b := heap.Pop(&bags).(Bag)
		for b.len < b.cap && additionalRocks > 0 {
			additionalRocks--
			b.len++
		}
		if b.len == b.cap {
			res++
		}
	}
	return res
}
