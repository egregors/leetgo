/*
	https://leetcode.com/problems/flatten-nested-list-iterator/

	You are given a nested list of integers nestedList. Each element is either an integer or a
	list whose elements may also be integers or other lists. Implement an iterator to flatten it.

	Implement the NestedIterator class:

		NestedIterator(List<NestedInteger> nestedList) Initializes the iterator with the nested
		list nestedList.
		int next() Returns the next integer in the nested list.
		boolean hasNext() Returns true if there are still some integers in the nested list and
		false otherwise.

	Your code will be tested with the following pseudocode:

	initialize iterator with nestedList
	res = []
	while iterator.hasNext()
		append iterator.next() to the end of res
	return res

	If res matches the expected flattened list, then your code will be judged as correct.
*/

//nolint:revive // it's ok
package solutions

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct {
	Val    int
	Nested []*NestedInteger
}

func NewNestedInteger() *NestedInteger {
	return &NestedInteger{
		Val:    -1,
		Nested: []*NestedInteger{},
	}
}

func (n NestedInteger) IsInteger() bool       { return len(n.Nested) == 0 }
func (n *NestedInteger) SetInteger(value int) { n.Val = value }
func (n NestedInteger) GetInteger() int {
	if n.IsInteger() {
		return n.Val
	}
	panic("not an Integer")
}
func (n *NestedInteger) Add(elem NestedInteger) {
	n.Nested = append(n.Nested, &elem)
}
func (n NestedInteger) GetList() []*NestedInteger {
	if !n.IsInteger() {
		return n.Nested
	}
	return []*NestedInteger{}
}

type NestedIterator struct {
	cache []int
	cur   int
}

// NewNestedIterator should call Constructor to pass LeetCode tests
func NewNestedIterator(nestedList []*NestedInteger) *NestedIterator {
	vals := &[]int{}
	for _, n := range nestedList {
		dfs341(n, vals)
	}
	return &NestedIterator{*vals, 0}
}

func dfs341(ns *NestedInteger, vals *[]int) {
	if ns.IsInteger() {
		*vals = append(*vals, ns.GetInteger())
		return
	}
	for _, n := range ns.GetList() {
		dfs341(n, vals)
	}
	return
}

func (i *NestedIterator) Next() int {
	if i.HasNext() {
		x := i.cache[0]
		i.cache = i.cache[1:]
		return x
	}
	return -1
}

func (i *NestedIterator) HasNext() bool {
	return len(i.cache) > 0
}
