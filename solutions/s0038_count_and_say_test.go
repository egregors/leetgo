package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{n: 1},
			want: "1",
		},
		{
			name: "Test 2",
			args: args{n: 2},
			want: "11",
		},
		{
			name: "Test 3",
			args: args{n: 3},
			want: "21",
		},
		{
			name: "Test 4",
			args: args{n: 4},
			want: "1211",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countAndSay(tt.args.n), "countAndSay(%v)", tt.args.n)
		})
	}
}
