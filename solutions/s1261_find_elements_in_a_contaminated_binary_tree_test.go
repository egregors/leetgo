package solutions

import (
	"testing"
)

func TestNewFind(t *testing.T) {
	tests := []struct {
		name    string
		tree    *TreeNode
		findSeq []int
		wantSeq []bool
	}{
		{
			name:    "testcase 1",
			tree:    NewTreeNode("[-1,null,-1]"),
			findSeq: []int{1, 2},
			wantSeq: []bool{false, true},
		},
		{
			name:    "testcase 2",
			tree:    NewTreeNode("[-1,-1,-1,-1,-1]"),
			findSeq: []int{1, 3, 5},
			wantSeq: []bool{true, true, false},
		},
		{
			name:    "testcase 3",
			tree:    NewTreeNode("[-1,null,-1,-1,null,-1]"),
			findSeq: []int{2, 3, 4, 5},
			wantSeq: []bool{true, false, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFindElement(tt.tree)
			for i, want := range tt.wantSeq {
				if got := f.Find(tt.findSeq[i]); got != want {
					t.Errorf("Find() = %v, want %v", got, want)
				}
			}
		})
	}
}
