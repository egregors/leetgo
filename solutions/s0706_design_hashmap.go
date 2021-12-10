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

const SPACE = 2069

type Pair struct {
	k, v int
}

type Bucket struct {
	b []Pair
}

func (b Bucket) Get(key int) int {
	for _, p := range b.b {
		if p.k == key {
			return p.v
		}
	}
	return -1
}

func (b *Bucket) Update(key, value int) {
	found := false
	for i, p := range b.b {
		if key == p.k {
			b.b[i] = Pair{key, value}
			found = true
			break
		}
	}

	if !found {
		b.b = append(b.b, Pair{key, value})
	}
}

func (b *Bucket) Remove(key int) {
	for i, p := range b.b {
		if p.k == key {
			b.b = append(b.b[:i], b.b[i+1:]...)
		}
	}
}

type MyHashMap struct {
	keySpace  int
	hashTable []Bucket
}

// NewMyHashMap is Constructor for MyHashMap (it must call Constructor for LeetCode)
func NewMyHashMap() MyHashMap {
	return MyHashMap{
		keySpace:  SPACE,
		hashTable: make([]Bucket, SPACE),
	}
}

func (m *MyHashMap) Put(key, value int) {
	hashKey := key % m.keySpace
	m.hashTable[hashKey].Update(key, value)
}

func (m MyHashMap) Get(key int) int {
	hashKey := key % m.keySpace
	return m.hashTable[hashKey].Get(key)
}

func (m *MyHashMap) Remove(key int) {
	hashKey := key % m.keySpace
	m.hashTable[hashKey].Remove(key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
