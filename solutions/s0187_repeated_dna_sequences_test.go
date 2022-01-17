package solutions

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	res := findRepeatedDnaSequences(s)
	sort.Strings(res)
	assert.Equal(t, []string{"AAAAACCCCC", "CCCCCAAAAA"}, res)
}
