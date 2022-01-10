package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCircularQueue(t *testing.T) {
	myCircularQueue := NewMyCircularQueue(3)
	assert.True(t, myCircularQueue.EnQueue(1))
	assert.True(t, myCircularQueue.EnQueue(2))
	assert.True(t, myCircularQueue.EnQueue(3))
	assert.False(t, myCircularQueue.EnQueue(4))
	assert.Equal(t, 3, myCircularQueue.Rear())
	assert.True(t, myCircularQueue.IsFull())
	assert.True(t, myCircularQueue.DeQueue())
	assert.True(t, myCircularQueue.EnQueue(4))
	assert.Equal(t, 4, myCircularQueue.Rear())

	myCircularQueue = NewMyCircularQueue(6)
	assert.True(t, myCircularQueue.EnQueue(6))
	assert.Equal(t, 6, myCircularQueue.Rear())
	assert.Equal(t, 6, myCircularQueue.Rear())
	assert.True(t, myCircularQueue.DeQueue())
	assert.True(t, myCircularQueue.EnQueue(5))
	assert.Equal(t, 5, myCircularQueue.Rear())
	assert.True(t, myCircularQueue.DeQueue())
	assert.Equal(t, -1, myCircularQueue.Front())
	assert.False(t, myCircularQueue.DeQueue())
	assert.False(t, myCircularQueue.DeQueue())
	assert.False(t, myCircularQueue.DeQueue())

}
