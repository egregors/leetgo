/*
	https://leetcode.com/problems/can-place-flowers/

	You have a long flowerbed in which some of the plots are planted, and some are not. However,
	flowers cannot be planted in adjacent plots.

	Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty,
	and an integer n, return if n new flowers can be planted in the flowerbed without violating the
	no-adjacent-flowers rule.
*/

package solutions

func canPlaceFlowers(flowerbed []int, n int) bool {
	i, count := 0, 0
	for i < len(flowerbed) {
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
		i++
	}
	return count >= n
}
