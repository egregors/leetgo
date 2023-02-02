/*
	https://leetcode.com/problems/lfu-cache/

	Design and implement a data structure for a Least Frequently Used (LFU) cache.

	Implement the LFUCache class:

		LFUCache(int capacity) Initializes the object with the capacity of the data structure.
		int get(int key) Gets the value of the key if the key exists in the cache. Otherwise,
	returns -1.
		void put(int key, int value) Update the value of the key if present, or inserts the key
	if not already present. When the cache reaches its capacity, it should invalidate and
	remove the least frequently used key before inserting a new item. For this problem, when
	there is a tie (i.e., two or more keys with the same frequency), the least recently used
	key would be invalidated.

	To determine the least frequently used key, a use counter is maintained for each key in
	the cache. The key with the smallest use counter is the least frequently used key.

	When a key is first inserted into the cache, its use counter is set to 1 (due to the put
	operation). The use counter for a key in the cache is incremented either a get or put
	operation is called on it.

	The functions get and put must each run in O(1) average time complexity.
*/

//nolint:revive // meh
package solutions

import "container/list"

type LFUCache struct {
	cap     int
	minF    int
	cache   map[int]*list.Element
	buckets map[int]*list.List
}

type LFUCacheElement struct {
	Key  int
	Val  int
	Freq int
}

// NewLFUCache should call Constructor to pass LeetCode tests
func NewLFUCache(capacity int) LFUCache {
	buckets := make(map[int]*list.List)
	buckets[1] = list.New()
	newCache := LFUCache{
		cap:     capacity,
		minF:    1,
		cache:   make(map[int]*list.Element),
		buckets: buckets,
	}
	return newCache
}

func (lfc *LFUCache) Get(key int) int {
	e, ok := lfc.cache[key]
	if ok {
		lfuCacheElement := e.Value.(*LFUCacheElement)
		oldBucket := lfc.buckets[lfuCacheElement.Freq]
		oldBucket.Remove(e)
		if lfuCacheElement.Freq == lfc.minF && oldBucket.Len() == 0 {
			lfc.minF++
		}
		lfuCacheElement.Freq++
		newBucket, ok := lfc.buckets[lfuCacheElement.Freq]
		if !ok {
			newBucket = list.New()
			lfc.buckets[lfuCacheElement.Freq] = newBucket
		}
		lfc.cache[key] = newBucket.PushFront(lfuCacheElement)
		return lfuCacheElement.Val
	} else {
		return -1
	}
}

func (lfc *LFUCache) Put(key int, value int) {
	e, ok := lfc.cache[key]
	if ok {
		lfuCacheElement := e.Value.(*LFUCacheElement)
		lfuCacheElement.Val = value
		lfc.Get(key)
	} else {
		newLfuCacheElement := LFUCacheElement{Key: key, Val: value, Freq: 1}
		newElement := lfc.buckets[1].PushFront(&newLfuCacheElement)
		lfc.cache[key] = newElement
		if len(lfc.cache) > lfc.cap {
			lfBucket := lfc.buckets[lfc.minF]
			lrElement := lfBucket.Back()
			lfBucket.Remove(lrElement)
			lfuCacheElement := lrElement.Value.(*LFUCacheElement)
			delete(lfc.cache, lfuCacheElement.Key)
		}
		lfc.minF = 1
	}
}
