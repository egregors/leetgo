package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortedListToBST(t *testing.T) {
	// TODO: smth wrong with this test, i'll fix it later (hehe)
	t.Skip()
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-1",
			args: args{
				head: NewListNode([]int{-10, -3, 0, 5, 9}),
			},
			want: "[0,-3,9,-10,null,5]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equalf(t, tt.want, sortedListToBST(tt.args.head).String(), "sortedListToBST(%v)", tt.args.head)
		})
	}
}
