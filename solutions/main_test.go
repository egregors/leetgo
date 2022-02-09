package solutions

import (
	"sort"
	"strconv"
	"strings"
)

// IsEqualAnyOrderInts returns true if both slices contain the same numbers independently of order
func IsEqualAnyOrderInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	aCp, bCp := make([]int, len(a)), make([]int, len(a))

	copy(aCp, a)
	copy(bCp, b)

	sort.Ints(aCp)
	sort.Ints(bCp)

	for i := 0; i < len(a); i++ {
		if aCp[i] != bCp[i] {
			return false
		}
	}
	return true
}

// IsEqualAnyOrderIntsSlices returns true if both slices of Int slices contain the same slices
// independently of order (don't use it for permutations!)
func IsEqualAnyOrderIntsSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	n := len(a)

	mA := make(map[string]bool)
	mB := make(map[string]bool)

	for i := 0; i < n; i++ {
		aCp := make([]int, len(a[i]))
		bCp := make([]int, len(b[i]))

		copy(aCp, a[i])
		copy(bCp, b[i])
		sort.Ints(aCp)
		sort.Ints(bCp)

		var aKey []string
		for _, x := range aCp {
			aKey = append(aKey, strconv.Itoa(x))
		}
		mA[strings.Join(aKey, ",")] = true

		var bKey []string
		for _, x := range bCp {
			bKey = append(bKey, strconv.Itoa(x))
		}
		mB[strings.Join(bKey, ",")] = true
	}

	if len(mB) > len(mA) {
		mA, mB = mB, mA
	}

	for k, v := range mA {
		if v2, ok := mB[k]; !ok {
			return false
		} else if v2 != v {
			return false
		}
	}

	return true
}
