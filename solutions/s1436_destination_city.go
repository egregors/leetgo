/*
	https://leetcode.com/problems/destination-city/description

	You are given the array paths, where paths[i] = [cityAi, cityBi] means there exists a direct path going from
	cityAi to cityBi. Return the destination city, that is, the city without any path outgoing to another city.

	It is guaranteed that the graph of paths forms a line without any loop, therefore, there will be
	exactly one destination city.
*/

package solutions

func destCity(paths [][]string) string {
	m := make(map[string]int)
	for _, p := range paths {
		m[p[0]]++
		m[p[1]]--
	}

	for k, v := range m {
		if v == -1 {
			return k
		}
	}

	return "FAIL"
}
