/*
	https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/description/

	You have information about n different recipes. You are given a string array
		recipes and a 2D
	string array ingredients. The ith recipe has the name recipes[i], and you can
		create it if you
	have all the needed ingredients from ingredients[i]. A recipe can also be an
		ingredient for other
	recipes, i.e., ingredients[i] may contain a string that is in recipes.

	You are also given a string array supplies containing all the ingredients that
		you initially have,
	and you have an infinite supply of all of them.

	Return a list of all the recipes that you can create. You may return the answer
		in any order.

	Note that two recipes may contain each other in their ingredients.
*/

package solutions

import "slices"

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	rsM := make(map[string]int)
	for i, r := range recipes {
		rsM[r] = i
	}

	supM := make(map[string]struct{})
	for _, s := range supplies {
		supM[s] = struct{}{}
	}

	var res []string
	for _, r := range recipes {
		if canDo(r, []string{}, ingredients, supM, rsM) {
			res = append(res, r)
		}
	}

	return res
}

func canDo(
	r string,
	prev []string,
	ings [][]string,
	supM map[string]struct{},
	rsM map[string]int,
) bool {
	iN, ok := rsM[r]
	if !ok {
		return false
	}

	ingsForR := ings[iN]

	for _, i := range ingsForR {
		if slices.Contains(prev, i) {
			return false
		} // loop
		if _, ok := supM[i]; !ok {
			prev = append(prev, r)
			if canDo(i, prev, ings, supM, rsM) {
				supM[i] = struct{}{}
			} else {
				return false
			}
		}
	}

	supM[r] = struct{}{}

	return true
}
