/*
	https://leetcode.com/problems/jump-game-iv

	Given an array of integers arr, you are initially positioned at the first index of the array.

	In one step you can jump from index i to index:

		i + 1 where: i + 1 < arr.length.
		i - 1 where: i - 1 >= 0.
		j where: arr[i] == arr[j] and i != j.

	Return the minimum number of steps to reach the last index of the array.

	Notice that you can not jump outside of the array at any time.
*/

package solutions

func minJumps(arr []int) int {
	jumps := make(map[int][]int)
	for i, v := range arr {
		jumps[v] = append(jumps[v], i)
	}

	steps := 0
	queue := []int{0}

	visited := make([]bool, len(arr))
	visited[0] = true

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == len(arr)-1 {
				return steps
			}
			for _, idx := range jumps[arr[cur]] {
				if !visited[idx] {
					visited[idx] = true
					queue = append(queue, idx)
				}
			}
			delete(jumps, arr[cur])
			if cur+1 < len(arr) && !visited[cur+1] {
				visited[cur+1] = true
				queue = append(queue, cur+1)
			}
			if cur-1 > -1 && !visited[cur-1] {
				visited[cur-1] = true
				queue = append(queue, cur-1)
			}
		}

		steps++
	}

	return -1
}
