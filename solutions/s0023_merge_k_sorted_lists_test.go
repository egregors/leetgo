package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				lists: []*ListNode{
					NewListNode([]int{1, 4, 5}),
					NewListNode([]int{1, 3, 4}),
					NewListNode([]int{2, 6}),
				},
			},
			want: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, mergeKLists(tt.args.lists).ToSlice(), "mergeKLists(%v)", tt.args.lists)
		})
	}
}
