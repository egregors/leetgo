package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countSubarrays3392(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Test1",
			nums: []int{1, 2, 1, 4, 1},
			want: 1,
		},
		{
			name: "Test2",
			nums: []int{1, 1, 1},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countSubarrays3392(tt.nums), "countSubarrays3392(%v)", tt.nums)
		})
	}
}
