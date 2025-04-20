package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestMonotonicSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		name string
		want int
	}{
		{
			name: "case 1",
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "case 2",
			nums: []int{1, 4, 3, 3, 2},
			want: 2,
		},
		{
			name: "case 3",
			nums: []int{3, 3, 3},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestMonotonicSubarray(tt.nums), "longestMonotonicSubarray(%v)", tt.nums)
		})
	}
}
