package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getHappyString(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testcase 1",
			args: args{1, 3},
			want: "c",
		},
		{
			name: "testcase 2",
			args: args{1, 4},
			want: "",
		},
		{
			name: "testcase 3",
			args: args{3, 9},
			want: "cab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getHappyString(tt.args.n, tt.args.k), "getHappyString(%v, %v)", tt.args.n, tt.args.k)
		})
	}
}
