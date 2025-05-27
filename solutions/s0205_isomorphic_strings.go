/*
	https://leetcode.com/problems/isomorphic-strings/description/

	Given two strings s and t, determine if they are isomorphic.

	Two strings s and t are isomorphic if the characters in s can be replaced to
		get t.

	All occurrences of a character must be replaced with another character while
		preserving
	the order of characters. No two characters may map to the same character, but a
		character
	may map to itself.
*/

package solutions

func isIsomorphic(s string, t string) bool {
	sm, tm := [128]byte{}, [128]byte{} // full ascii -> 256

	for i := 0; i < len(s); i++ {
		if sm[s[i]] != tm[t[i]] {
			return false
		}
		sm[s[i]] = byte(i + 1)
		tm[t[i]] = byte(i + 1)
	}

	return true
}

func isIsomorphic1(s string, t string) bool { //nolint:unused // it's ok
	if len(s) != len(t) {
		return false
	}

	s2t, t2s := make(map[byte]byte), make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sCh, tCh := s[i], t[i]
		if fromT, ok := s2t[sCh]; !ok {
			if _, ok := t2s[tCh]; ok {
				return false
			}
			s2t[sCh] = tCh
			t2s[tCh] = sCh
		} else {
			if fromT != tCh {
				return false
			}
		}
	}

	return true
}

func isIsomorphic2(s string, t string) bool { //nolint:unused // it's ok
	if len(s) != len(t) {
		return false
	}

	sP := make([]int, len(s))
	tP := make([]int, len(t))
	sM := make(map[byte]int)
	tM := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		// s
		if _, ok := sM[s[i]]; !ok {
			sM[s[i]] = i
		}
		sP[i] = sM[s[i]]

		// t
		if _, ok := tM[t[i]]; !ok {
			tM[t[i]] = i
		}
		tP[i] = tM[t[i]]

	}

	for i, e := range sP {
		if tP[i] != e {
			return false
		}
	}

	return true
}
