package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_triangularSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test case 1",
			nums: []int{1, 2, 3, 4, 5},
			want: 8,
		},
		{
			name: "test case 2",
			nums: []int{5},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, triangularSum(tt.nums), "triangularSum(%v)", tt.nums)
		})
	}
}
