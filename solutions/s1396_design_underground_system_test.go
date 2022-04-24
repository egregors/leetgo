package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndergroundSystem1(t *testing.T) {
	undergroundSystem := NewUndergroundSystem()
	undergroundSystem.CheckIn(45, "Leyton", 3)
	undergroundSystem.CheckIn(32, "Paradise", 8)
	undergroundSystem.CheckIn(27, "Leyton", 10)
	undergroundSystem.CheckOut(45, "Waterloo", 15)
	undergroundSystem.CheckOut(27, "Waterloo", 20)
	undergroundSystem.CheckOut(32, "Cambridge", 22)
	assert.Equal(t, float64(14), undergroundSystem.GetAverageTime("Paradise", "Cambridge"))
	assert.Equal(t, float64(11), undergroundSystem.GetAverageTime("Leyton", "Waterloo"))
	undergroundSystem.CheckIn(10, "Leyton", 24)
	assert.Equal(t, float64(11), undergroundSystem.GetAverageTime("Leyton", "Waterloo"))
	undergroundSystem.CheckOut(10, "Waterloo", 38)
	assert.Equal(t, float64(12), undergroundSystem.GetAverageTime("Leyton", "Waterloo"))
}

func TestUndergroundSystem2(t *testing.T) {
	undergroundSystem := NewUndergroundSystem()
	undergroundSystem.CheckIn(10, "Leyton", 3)
	undergroundSystem.CheckOut(10, "Paradise", 8)
	assert.Equal(t, float64(5), undergroundSystem.GetAverageTime("Leyton", "Paradise"))
	undergroundSystem.CheckIn(5, "Leyton", 10)
	undergroundSystem.CheckOut(5, "Paradise", 16)
	assert.Equal(t, 5.5, undergroundSystem.GetAverageTime("Leyton", "Paradise"))
	undergroundSystem.CheckIn(2, "Leyton", 21)
	undergroundSystem.CheckOut(2, "Paradise", 30)
	assert.Equal(t, 6.666666666666667, undergroundSystem.GetAverageTime("Leyton", "Paradise"))
}
