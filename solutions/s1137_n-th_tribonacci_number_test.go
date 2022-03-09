package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tribonacci(t *testing.T) {
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
				n: 4,
			},
			want: 4,
		},
		{
			name: "Test 2",
			args: args{
				n: 25,
			},
			want: 1389537,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tribonacci(tt.args.n), "tribonacci(%v)", tt.args.n)
		})
	}
}
