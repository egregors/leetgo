package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{path: "/home/"},
			want: "/home",
		},
		{
			name: "Test Case 2",
			args: args{path: "/../"},
			want: "/",
		},
		{
			name: "Test Case 3",
			args: args{path: "/home//foo/"},
			want: "/home/foo",
		},
		{
			name: "Test Case 4",
			args: args{path: "/.../a/../b/c/../d/./"},
			want: "/.../b/d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, simplifyPath(tt.args.path), "simplifyPath(%v)", tt.args.path)
		})
	}
}
