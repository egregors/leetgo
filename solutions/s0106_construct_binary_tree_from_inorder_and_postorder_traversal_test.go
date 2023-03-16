package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildTree106(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{
				[]int{9, 3, 15, 20, 7},
				[]int{9, 15, 7, 20, 3},
			},
			"[3,9,20,null,null,15,7]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, buildTree106(tt.args.inorder, tt.args.postorder).String(), "buildTree106(%v, %v)", tt.args.inorder, tt.args.postorder)
		})
	}
}
