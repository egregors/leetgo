package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			want: -1,
		},
		{
			name: "Example 3",
			args: args{
				haystack: "aaaaa",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{
				haystack: "",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "Example 5",
			args: args{
				haystack: "",
				needle:   "a",
			},
			want: -1,
		},
		{
			name: "Example 6",
			args: args{
				haystack: "mississippi",
				needle:   "issipi",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, strStr(tt.args.haystack, tt.args.needle), "strStr(%v, %v)", tt.args.haystack, tt.args.needle)
		})
	}
}
