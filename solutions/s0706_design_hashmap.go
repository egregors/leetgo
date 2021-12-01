/*
	https://leetcode.com/problems/design-hashmap/

	Design a HashMap without using any built-in hash table libraries.

	Implement the MyHashMap class:

		MyHashMap() initializes the object with an empty map.
		void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already
		exists in the map, update the corresponding value.
		int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains
		no mapping for the key.
		void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.
*/

//nolint:revive //it's ok
package solutions

type pair struct {
	fst, snd int
}

type MyHashMap struct {
	xs []pair
}

// NewMyHashMap is Constructor for MyHashMap (it must call Constructor for LeetCode)
func NewMyHashMap() MyHashMap {
	return MyHashMap{xs: []pair{}}
}

func (m *MyHashMap) Put(key, value int) {
	if m.Get(key) != -1 {
		m.Remove(key)
	}
	m.xs = append(m.xs, pair{key, value})
}

func (m MyHashMap) Get(key int) int {
	for _, p := range m.xs {
		if p.fst == key {
			return p.snd
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	for i, p := range m.xs {
		if p.fst == key {
			m.xs = append(m.xs[:i], m.xs[i+1:]...)
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
