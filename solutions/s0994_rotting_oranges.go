/*
	https://leetcode.com/problems/rotting-oranges/

	You are given an m x n grid where each cell can have one of three values:

		0 representing an empty cell,
		1 representing a fresh orange, or
		2 representing a rotten orange.

	Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange
	becomes rotten.

	Return the minimum number of minutes that must elapse until no cell has a fresh orange.
	If this is impossible, return -1.
*/

//nolint:revive //it's ok
package solutions

import "fmt"

type Orange struct {
	r, c int
}
type Queue994 struct {
	xs []Orange
}

func (q Queue994) IsEmpty() bool {
	return len(q.xs) == 0
}

func (q *Queue994) PushRight(o Orange) {
	q.xs = append(q.xs, o)
}

func (q *Queue994) PopLeft() Orange {
	o := q.xs[0]
	q.xs = q.xs[1:]
	return o
}

func orangesRotting(grid [][]int) int {
	queue := Queue994{xs: []Orange{}}

	freshOranges := 0
	rows, cols := len(grid), len(grid[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue.PushRight(Orange{r, c})
			} else if grid[r][c] == 1 {
				freshOranges++
			}
		}
	}

	fmt.Println(queue.xs)

	queue.PushRight(Orange{-1, -1})

	mins := -1
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for !queue.IsEmpty() {
		o := queue.PopLeft()
		if o.r == -1 {
			mins++
			if !queue.IsEmpty() {
				queue.PushRight(Orange{-1, -1})
			}
		} else {
			for _, d := range directions {
				neighborRow := o.r + d[0]
				neighborCol := o.c + d[1]
				if neighborRow >= 0 && neighborRow < rows && neighborCol >= 0 && neighborCol < cols {
					if grid[neighborRow][neighborCol] == 1 {
						grid[neighborRow][neighborCol] = 2
						freshOranges--
						queue.PushRight(Orange{neighborRow, neighborCol})
					}
				}
			}
		}
	}
	if freshOranges == 0 {
		return mins
	}
	return -1
}
