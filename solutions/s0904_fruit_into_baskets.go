/*
	https://leetcode.com/problems/fruit-into-baskets/

	You are visiting a farm that has a single row of fruit trees arranged from left
		to right. The trees are
	represented by an integer array fruits where fruits[i] is the type of fruit the
		ith tree produces.

You want to collect as much fruit as possible. However, the owner has some
	strict rules that you must follow:

    You only have two baskets, and each basket can only hold a single type of
    	fruit. There is no limit on the amount
	of fruit each basket can hold.
    Starting from any tree of your choice, you must pick exactly one fruit from
    	every tree (including the start tree)
	while moving to the right. The picked fruits must fit in one of your baskets.
    Once you reach a tree with fruit that cannot fit in your baskets, you must
    	stop.

Given the integer array fruits, return the maximum number of fruits you can
	pick.
*/

package solutions

func totalFruit(fruits []int) int {
	m := make(map[int]int, 3)
	var maxfruits int

	for i, j := 0, 0; i < len(fruits); i++ {
		m[fruits[i]]++
		for len(m) > 2 {
			m[fruits[j]]--
			if m[fruits[j]] == 0 {
				delete(m, fruits[j])
			}
			j++
		}
		if i-j+1 > maxfruits {
			maxfruits = i - j + 1
		}
	}

	return maxfruits
}
