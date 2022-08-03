package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCalendar(t *testing.T) {
	c := NewMyCalendar()
	assert.True(t, c.Book(10, 20))
	assert.False(t, c.Book(15, 25))
	assert.True(t, c.Book(20, 30))
}
