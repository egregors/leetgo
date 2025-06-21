package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isStrobogrammatic(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{num: "69"},
			want: true,
		},
		{
			name: "Example 2",
			args: args{num: "88"},
			want: true,
		},
		{
			name: "Example 3",
			args: args{num: "962"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isStrobogrammatic(tt.args.num), "isStrobogrammatic(%v)", tt.args.num)
		})
	}
}
