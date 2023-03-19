package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBrowserHistory(t *testing.T) {
	bh := NewBrowserHistory("leetcode.com")
	bh.Visit("google.com")
	bh.Visit("facebook.com")
	bh.Visit("youtube.com")
	assert.Equal(t, "facebook.com", bh.Back(1))
	assert.Equal(t, "google.com", bh.Back(1))
	assert.Equal(t, "facebook.com", bh.Forward(1))
	bh.Visit("linkedin.com")
	assert.Equal(t, "linkedin.com", bh.Forward(2))
	assert.Equal(t, "google.com", bh.Back(2))
	assert.Equal(t, "leetcode.com", bh.Back(7))
}
