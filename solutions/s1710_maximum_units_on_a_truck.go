/*
	https://leetcode.com/problems/maximum-units-on-a-truck/

	You are assigned to put some amount of boxes onto one truck. You are given a 2D array boxTypes,
	where boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi]:

		numberOfBoxesi is the number of boxes of type i.
		numberOfUnitsPerBoxi is the number of units in each box of the type i.

	You are also given an integer truckSize, which is the maximum number of boxes that can be put on
	the truck. You can choose any boxes to put on the truck as long as the number of boxes does not exceed truckSize.

	Return the maximum total number of units that can be put on the truck.
*/

package solutions

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	ans := 0
	for _, t := range boxTypes {
		for t[0] > 0 {
			ans += t[1]
			t[0]--
			truckSize--
			if truckSize == 0 {
				return ans
			}
		}
	}

	return ans
}
