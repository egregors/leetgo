package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_modifiedList(t *testing.T) {
	type args struct {
		nums []int
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "testcase 1",
			args: args{
				nums: []int{1, 2, 3},
				head: NewListNode([]int{3, 2, 1, 4, 5}),
			},
			want: NewListNode([]int{4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, modifiedList(tt.args.nums, tt.args.head), "modifiedList(%v, %v)", tt.args.nums, tt.args.head)
		})
	}
}
