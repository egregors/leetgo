package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyLinkedList(t *testing.T) {
	// Case 1

	// "MyLinkedList",   	[]
	// "addAtHead",   		[7]
	// "addAtHead",   		[2]
	// "addAtHead",   		[1]
	// "addAtIndex",   		[3,0]
	// "deleteAtIndex",		[2]
	// "addAtHead",   		[6]
	// "addAtTail",   		[4]
	// "get",   			[4]
	// "addAtHead",   		[4]
	// "addAtIndex",   		[5,0]
	// "addAtHead",   		[6]

	list := NewMyLinkedList()
	list.AddAtHead(7)
	list.AddAtHead(2)
	list.AddAtHead(1)
	list.AddAtIndex(3, 0)
	list.DeleteAtIndex(2)
	list.AddAtHead(6)
	list.AddAtTail(4)
	assert.Equal(t, 4, list.Get(4))
	list.AddAtHead(4)
	list.AddAtIndex(5, 0)
	list.AddAtHead(6)

	// Case 2
	// "MyLinkedList",		[],
	// "addAtIndex",		[0, 10],
	// "addAtIndex",		[0, 20],
	// "addAtIndex",		[1, 30],
	// "get",				[0]

	list = NewMyLinkedList()
	list.AddAtIndex(0, 10)
	list.AddAtIndex(0, 20)
	list.AddAtIndex(1, 30)
	assert.Equal(t, 20, list.Get(0))

	// Case 3
	// "MyLinkedList",	[],
	// "addAtTail",		[1],
	// "get", 			[0]

	list = NewMyLinkedList()
	list.AddAtTail(1)
	assert.Equal(t, 1, list.Get(0))
}
