package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_connect(t *testing.T) {
	n4 := &Node{
		Val:   4,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	n5 := &Node{
		Val:   5,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	n6 := &Node{
		Val:   6,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	n7 := &Node{
		Val:   7,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}

	n2 := &Node{
		Val:   2,
		Left:  n4,
		Right: n5,
		Next:  nil,
	}
	n3 := &Node{
		Val:   3,
		Left:  n6,
		Right: n7,
		Next:  nil,
	}

	n1 := &Node{
		Val:   1,
		Left:  n2,
		Right: n3,
		Next:  nil,
	}

	connect(n1)

	assert.Nil(t, n1.Next)
	assert.Equal(t, n3, n2.Next)
	assert.Nil(t, n3.Next)
	assert.Equal(t, n5, n4.Next)
	assert.Equal(t, n6, n5.Next)
	assert.Equal(t, n7, n6.Next)
	assert.Nil(t, n7.Next)
}
