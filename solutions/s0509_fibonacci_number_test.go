package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "Test 3",
			args: args{
				n: 4,
			},
			want: 3,
		},
		{
			name: "Test 4",
			args: args{
				n: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, fib(tt.args.n), "fib(%v)", tt.args.n)
		})
	}
}
