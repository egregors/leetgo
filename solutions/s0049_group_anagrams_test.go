package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_groupAnagrams(t *testing.T) {
	// t1
	assert.Equal(t, [][]string{
		{"", ""},
	}, groupAnagrams([]string{"", ""}))

	// t2
	r := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	for _, el := range r {
		switch len(el) {
		case 1:
			assert.Equal(t, el, []string{"bat"})
		case 2:
			assert.Equal(t, el, []string{"tan", "nat"})
		case 3:
			assert.Equal(t, el, []string{"eat", "tea", "ate"})
		}
	}
}
