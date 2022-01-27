/*
	https://leetcode.com/problems/keys-and-rooms/

	There are n rooms labeled from 0 to n - 1 and all the rooms are locked except for room 0.
	Your goal is to visit all the rooms. However, you cannot enter a locked room without having its key.

	When you visit a room, you may find a set of distinct keys in it. Each key has a number on it,
	denoting which room it unlocks, and you can take all of them with you to unlock the other rooms.

	Given an array rooms where rooms[i] is the set of keys that you can obtain if you visited room i,
	return true if you can visit all the rooms, or false otherwise.
*/

package solutions

func visit(room []int, rooms [][]int, visited map[int]bool) {
	for _, r := range room {
		if !visited[r] {
			visited[r] = true
			visit(rooms[r], rooms, visited)
		}
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	visited := map[int]bool{0: true}
	visit(rooms[0], rooms, visited)
	return len(visited) == len(rooms)
}
