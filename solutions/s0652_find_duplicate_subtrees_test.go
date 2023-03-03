package solutions

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_findDuplicateSubtrees(t *testing.T) {
	root := NewTreeNode("[1,2,3,4,null,2,4,null,null,4]")
	want := "[[4],[2,4]]"
	assert.Equal(t, want, func() string {
		ns := findDuplicateSubtrees(root)
		var res []string
		for _, n := range ns {
			res = append(res, n.String())
		}
		return "[" + strings.Join(res, ",") + "]"
	}())
}
