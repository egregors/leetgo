/*
	https://leetcode.com/problems/design-hashset/

	Design a HashSet without using any built-in hash table libraries.

	Implement MyHashSet class:

		void add(key) Inserts the value key into the HashSet.
		bool contains(key) Returns whether the value key exists in the HashSet or not.
		void remove(key) Removes the value key in the HashSet. If key does not exist
			in the HashSet, do nothing.
*/

//nolint:revive // it's ok
package solutions

const space = 2069

type bucket []int

func (b *bucket) Add(key int) {
	if !b.Contains(key) {
		*b = append(*b, key)
	}
}

func (b *bucket) Remove(key int) {
	for i := 0; i < len(*b); i++ {
		if (*b)[i] == key {
			*b = append((*b)[:i], (*b)[i+1:]...)
		}
	}
}

func (b *bucket) Contains(key int) bool {
	for _, k := range *b {
		if k == key {
			return true
		}
	}
	return false
}

type MyHashSet struct {
	keySpace  int
	hashTable []bucket
}

// NewMyHashSet should call Constructor to pass LeetCode tests
func NewMyHashSet() MyHashSet {
	return MyHashSet{
		keySpace:  space,
		hashTable: make([]bucket, space),
	}
}

func (s *MyHashSet) Add(key int) {
	hashKey := key % s.keySpace
	s.hashTable[hashKey].Add(key)
}

func (s *MyHashSet) Remove(key int) {
	hashKey := key % s.keySpace
	s.hashTable[hashKey].Remove(key)
}

func (s *MyHashSet) Contains(key int) bool {
	hashKey := key % s.keySpace
	return s.hashTable[hashKey].Contains(key)
}
