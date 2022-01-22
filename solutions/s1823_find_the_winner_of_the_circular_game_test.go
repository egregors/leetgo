package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findTheWinner(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				n: 5,
				k: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTheWinner(tt.args.n, tt.args.k), "findTheWinner(%v, %v)", tt.args.n, tt.args.k)
		})
	}
}
