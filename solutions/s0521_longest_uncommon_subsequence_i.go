package solutions

func findLUSlength(a string, b string) int { //nolint:gocritic // it's ok
	if a == b {
		return -1
	}
	return Maximum(len(a), len(b))
}
