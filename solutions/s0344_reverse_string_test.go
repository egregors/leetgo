package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseString(t *testing.T) {
	s := []byte("hello")
	reverseString(s)
	assert.Equal(t, []byte("olleh"), s)
}
