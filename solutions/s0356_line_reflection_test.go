package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isReflected(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   bool
	}{
		{
			name:   "Test 1",
			points: [][]int{{1, 1}, {-1, 1}},
			want:   true,
		},
		{
			name:   "Test 2",
			points: [][]int{{1, 1}, {-1, -1}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isReflected(tt.points), "isReflected(%v)", tt.points)
		})
	}
}
