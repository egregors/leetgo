package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Example 1",
			args: args{
				root: NewTreeNode("[3,9,20,null,null,15,7]"),
			},
			want: []float64{3.00000, 14.50000, 11.00000},
		},
		{
			name: "Example 1",
			args: args{
				root: NewTreeNode("[3,9,20,15,7]"),
			},
			want: []float64{3.00000, 14.50000, 11.00000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, averageOfLevels(tt.args.root), "averageOfLevels(%v)", tt.args.root)
		})
	}
}
