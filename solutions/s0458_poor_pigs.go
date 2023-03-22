/*
	https://leetcode.com/problems/poor-pigs/

	There are buckets buckets of liquid, where exactly one of the buckets is poisonous. To figure out
	which one is poisonous, you feed some number of (poor) pigs the liquid to see whether they will
	die or not. Unfortunately, you only have minutesToTest minutes to determine which bucket is poisonous.

	You can feed the pigs according to these steps:

		Choose some live pigs to feed.
		For each pig, choose which buckets to feed it. The pig will consume all the chosen buckets
		simultaneously and will take no time.
		Wait for minutesToDie minutes. You may not feed any other pigs during this time.
		After minutesToDie minutes have passed, any pigs that have been fed the poisonous bucket will die,
		and all others will survive.
		Repeat this process until you run out of time.

	Given buckets, minutesToDie, and minutesToTest, return the minimum number of pigs needed to figure out
		which bucket is poisonous within the allotted time.
*/

package solutions

import "math"

func poorPigs(buckets, minutesToDie, minutesToTest int) int {
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(minutesToTest/minutesToDie+1))))
}
