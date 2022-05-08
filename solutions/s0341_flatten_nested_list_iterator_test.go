package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNestedIterator(t *testing.T) {
	root := NewNestedInteger()
	n1 := NewNestedInteger()
	n11 := NewNestedInteger()
	n12 := NewNestedInteger()
	n2 := NewNestedInteger()
	n3 := NewNestedInteger()
	n31 := NewNestedInteger()
	n32 := NewNestedInteger()

	n11.SetInteger(1)
	n1.Add(*n11)
	n12.SetInteger(1)
	n1.Add(*n12)

	n2.SetInteger(2)

	n31.SetInteger(1)
	n3.Add(*n31)
	n32.SetInteger(1)
	n3.Add(*n32)

	root.Add(*n1)
	root.Add(*n2)
	root.Add(*n3)

	iter := NewNestedIterator(root.GetList())
	var res []int

	for iter.HasNext() {
		res = append(res, iter.Next())
	}
	assert.Equal(t, []int{1, 1, 2, 1, 1}, res)
}
