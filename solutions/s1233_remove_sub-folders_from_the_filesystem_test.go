package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeSubfolders(t *testing.T) {
	type args struct {
		folder []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				folder: []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"},
			},
			want: []string{"/a", "/c/d", "/c/f"},
		},
		{
			name: "Example 2",
			args: args{
				folder: []string{"/a", "/a/b/c", "/a/b/d"},
			},
			want: []string{"/a"},
		},
		{
			name: "Example 3",
			args: args{
				folder: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
			},
			want: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeSubfolders(tt.args.folder), "removeSubfolders(%v)", tt.args.folder)
		})
	}
}
