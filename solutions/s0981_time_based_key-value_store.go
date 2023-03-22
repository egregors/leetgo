/*
	https://leetcode.com/problems/time-based-key-value-store/

	Design a time-based key-value data structure that can store multiple values for the
	same key at different time stamps and retrieve the key's value at a certain timestamp.

	Implement the TimeMap class:

		TimeMap() Initializes the object of the data structure.
		void set(String key, String value, int timestamp) Stores the key key with the
			value value at the given time timestamp.
		String get(String key, int timestamp) Returns a value such that set was called
			previously, with timestamp_prev <= timestamp. If there are multiple such values,
			it returns the value associated with the largest timestamp_prev. If there are no
			values, it returns "".
*/

//nolint:revive // meh
package solutions

type TMPair struct {
	timestamp int
	value     string
}

type Coll []TMPair

func (c *Coll) Add(p TMPair) {
	if len(*c) == 0 {
		*c = append(*c, p)
		return
	}

	if p.timestamp <= (*c)[0].timestamp {
		*c = append(Coll{p}, *c...)
		return
	}

	if p.timestamp >= (*c)[len(*c)-1].timestamp {
		*c = append(*c, p)
		return
	}

	i := indexBs(*c, p.timestamp)
	if (*c)[i].timestamp > p.timestamp {
		i--
	}

	*c = append((*c)[:i], append(Coll{p}, (*c)[:i+1]...)...)
}

func (c Coll) Get(ts int) TMPair {
	i := indexBs(c, ts)
	if c[i].timestamp > ts {
		return c[i-1]
	}
	return c[i]
}

func indexBs(c Coll, ts int) int {
	lo, hi := 0, len(c)-1
	var mid int
	for lo <= hi {
		mid = (lo + hi) / 2
		if ts == c[mid].timestamp {
			return mid
		}

		if ts < c[mid].timestamp {
			hi = mid - 1
		}

		if ts > c[mid].timestamp {
			lo = mid + 1
		}
	}
	return mid
}

type TimeMap struct {
	xs map[string]*Coll
}

// NewTimeMap should call Constructor to pass LeetCode tests
func NewTimeMap() TimeMap {
	return TimeMap{
		xs: make(map[string]*Coll),
	}
}

func (tm *TimeMap) Set(key, value string, timestamp int) {
	if _, ok := tm.xs[key]; !ok {
		tm.xs[key] = &Coll{}
	}
	tm.xs[key].Add(TMPair{timestamp: timestamp, value: value})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	if _, ok := tm.xs[key]; !ok {
		return ""
	}
	if timestamp < (*tm.xs[key])[0].timestamp {
		return ""
	}
	return tm.xs[key].Get(timestamp).value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
