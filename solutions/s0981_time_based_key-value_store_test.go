package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimeMap1(t *testing.T) {
	timeMap := NewTimeMap()
	timeMap.Set("foo", "bar", 1)                   // store the key "foo" and value "bar" along with timestamp = 1.
	assert.Equal(t, "bar", timeMap.Get("foo", 1))  // return "bar"
	assert.Equal(t, "bar", timeMap.Get("foo", 3))  // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
	timeMap.Set("foo", "bar2", 4)                  // store the key "foo" and value "bar2" along with timestamp = 4.
	assert.Equal(t, "bar2", timeMap.Get("foo", 4)) // return "bar2"
	assert.Equal(t, "bar2", timeMap.Get("foo", 5)) // return "bar2"
}
func TestTimeMap2(t *testing.T) {
	timeMap := NewTimeMap()
	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)
	assert.Equal(t, "", timeMap.Get("love", 5))
	assert.Equal(t, "high", timeMap.Get("love", 10))
	assert.Equal(t, "high", timeMap.Get("love", 15))
	assert.Equal(t, "low", timeMap.Get("love", 20))
	assert.Equal(t, "low", timeMap.Get("love", 25))
}
