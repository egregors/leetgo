/*
	https://leetcode.com/problems/peeking-iterator/

		Design an iterator that supports the peek operation on an existing iterator in
			addition to the hasNext and the next operations.

	Implement the PeekingIterator class:

		PeekingIterator(Iterator<int> nums) Initializes the object with the given
			integer
		iterator iterator.
		int next() Returns the next element in the array and moves the pointer to the
			next element.
		boolean hasNext() Returns true if there are still elements in the array.
		int peek() Returns the next element in the array without moving the pointer.

	Note: Each language may have a different implementation of the constructor and
		Iterator, but
	they all support the int next() and boolean hasNext() functions.
*/

//nolint:revive // it's ok
package solutions

type Iterator interface {
	hasNext() bool
	next() int
}
type PeekingIterator struct {
	iter Iterator
	p    int
}

func NewPeekingIterator(iter Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter: iter,
		p:    -1,
	}
}

func (i *PeekingIterator) hasNext() bool {
	return i.p != -1 || i.iter.hasNext()
}

func (i *PeekingIterator) next() int {
	if i.p != -1 {
		x := i.p
		i.p = -1
		return x
	}
	if !i.iter.hasNext() {
		return -1
	}
	return i.iter.next()
}

func (i *PeekingIterator) peek() int {
	if i.p == -1 {
		if !i.iter.hasNext() {
			return -1
		}
		i.p = i.iter.next()
	}
	return i.p
}
