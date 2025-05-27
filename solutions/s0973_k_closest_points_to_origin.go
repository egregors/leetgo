/*
	https://leetcode.com/problems/k-closest-points-to-origin/

	Given an array of points where points[i] = [xi, yi] represents a point on the
		X-Y plane and an integer k,
	return the k closest points to the origin (0, 0).

	The distance between two points on the X-Y plane is the Euclidean distance
		(i.e., √(x1 - x2)2 + (y1 - y2)2).

	You may return the answer in any order. The answer is guaranteed to be unique
		(except for the order that it is in).
*/

//nolint:revive // it's ok
package solutions

import (
	"container/heap"
	"math"
)

type Point struct {
	X, Y int
	Dist float64
}

type pointHeap []*Point

func (h pointHeap) Len() int           { return len(h) }
func (h pointHeap) Less(i, j int) bool { return h[i].Dist > h[j].Dist }
func (h pointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *pointHeap) Push(x interface{}) { *h = append(*h, x.(*Point)) }

func (h *pointHeap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func getDist(x, y int) float64 {
	fx, fy := float64(x), float64(y)
	return math.Sqrt(fx*fx + fy*fy)
}

func kClosest(points [][]int, k int) [][]int {
	h := &pointHeap{}
	for _, p := range points {
		x, y := p[0], p[1]
		heap.Push(h, &Point{X: x, Y: y, Dist: getDist(x, y)})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	var res [][]int
	for h.Len() > 0 {
		p := heap.Pop(h).(*Point)
		res = append(res, []int{p.X, p.Y})
	}
	return res
}
