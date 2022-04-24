/*
	https://leetcode.com/problems/design-underground-system/

	An underground railway system is keeping track of customer travel times between different stations.
	They are using this data to calculate the average time it takes to travel from one station to another.

	Implement the UndergroundSystem class:

		void checkIn(int id, string stationName, int t)
			A customer with a card ID equal to id, checks in at the station stationName at time t.
			A customer can only be checked into one place at a time.
		void checkOut(int id, string stationName, int t)
			A customer with a card ID equal to id, checks out from the station stationName at time t.
		double getAverageTime(string startStation, string endStation)
			Returns the average time it takes to travel from startStation to endStation.
			The average time is computed from all the previous traveling times from startStation to endStation
			that happened directly, meaning a check in at startStation followed by a check out from endStation.
			The time it takes to travel from startStation to endStation may be different from the time it takes to
			travel from endStation to startStation.
			There will be at least one customer that has traveled from startStation to endStation before getAverageTime
			is called.

	You may assume all calls to the checkIn and checkOut methods are consistent.
	If a customer checks in at time t1 then checks out at time t2, then t1 < t2.
	All events happen in chronological order.
*/

//nolint:revive // it's ok
package solutions

import "fmt"

type Ride struct {
	id, in, out int
	from, to    string
}

type UndergroundSystem struct {
	openRides map[int]Ride
	stats     map[string][]int
}

// NewUndergroundSystem should call Constructor() to pass LeetCode tests
func NewUndergroundSystem() UndergroundSystem {
	return UndergroundSystem{
		openRides: make(map[int]Ride),
		stats:     make(map[string][]int),
	}
}

func (s *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	s.openRides[id] = Ride{
		id:   id,
		in:   t,
		from: stationName,
	}
}

func (s *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if r, ok := s.openRides[id]; ok {
		r.to = stationName
		r.out = t

		key := genKey(r.from, r.to)
		s.stats[key] = append(s.stats[key], r.out-r.in)
		delete(s.openRides, id)
	}
}

func (s *UndergroundSystem) GetAverageTime(startStation, endStation string) float64 {
	rs := s.stats[genKey(startStation, endStation)]
	return float64(sum(rs...)) / float64(len(rs))
}

func genKey(s1, s2 string) string {
	return fmt.Sprintf("%s_%s", s1, s2)
}

func sum(xs ...int) int {
	s := 0
	for _, x := range xs {
		s += x
	}
	return s
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
