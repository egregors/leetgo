package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getDecimalValue(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want int
	}{
		{
			name: "Test 1",
			head: NewListNode([]int{1, 0, 1}),
			want: 5,
		},
		{
			name: "Test 2",
			head: NewListNode([]int{0}),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getDecimalValue(tt.head), "getDecimalValue(%v)", tt.head)
		})
	}
}
