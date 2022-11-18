package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isUgly(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{6},
			want: true,
		},
		{
			name: "test2",
			args: args{8},
			want: true,
		},
		{
			name: "test3",
			args: args{14},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isUgly(tt.args.n), "isUgly(%v)", tt.args.n)
		})
	}
}
