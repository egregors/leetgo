package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_levelOrder429(t *testing.T) {
	t.Skip()
	// todo: i wanna make a tool to build n-arn trees, like NewTreeNode for bTree
	//  see: https://github.com/egregors/TreeNode/issues/7
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// todo: add tests here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, levelOrder429(tt.args.root), "levelOrder429(%v)", tt.args.root)
		})
	}
}
