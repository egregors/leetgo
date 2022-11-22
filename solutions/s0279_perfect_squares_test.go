package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numSquares(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				n: 12,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				n: 13,
			},
			want: 2,
		},
		{
			name: "test3",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "test4",
			args: args{
				n: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numSquares(tt.args.n), "numSquares(%v)", tt.args.n)
		})
	}
}
