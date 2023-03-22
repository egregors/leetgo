/*
	https://leetcode.com/problems/mirror-reflection/

	There is a special square room with mirrors on each of the four walls. Except for the
	southwest corner, there are receptors on each of the remaining corners, numbered 0, 1, and 2.

	The square room has walls of length p and a laser ray from the southwest corner first meets
	the east wall at a distance q from the 0th receptor.

	Given the two integers p and q, return the number of the receptor that the ray meets first.

	The test cases are guaranteed so that the ray will meet a receptor eventually.
*/

//nolint:revive // it's ok
package solutions

func mirrorReflection(p, q int) int {
	lcm := LCM(p, q)
	h := lcm / p
	w := lcm / q
	if isOdd(h) {
		if isOdd(w) {
			return 1
		}
		return 2
	}
	return 0
}

func GCD(p, q int) int {
	for q != 0 {
		r := p % q
		p = q
		q = r
	}
	return p
}

func LCM(p, q int) int {
	gcd := GCD(p, q)
	return p * q / gcd
}

func isOdd(n int) bool {
	return n%2 == 1
}
