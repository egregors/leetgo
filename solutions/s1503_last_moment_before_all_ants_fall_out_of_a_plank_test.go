package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getLastMoment(t *testing.T) {
	type args struct {
		n     int
		left  []int
		right []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n:     4,
				left:  []int{4, 3},
				right: []int{0, 1},
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				n:     7,
				left:  []int{},
				right: []int{0, 1, 2, 3, 4, 5, 6, 7},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getLastMoment(tt.args.n, tt.args.left, tt.args.right), "getLastMoment(%v, %v, %v)", tt.args.n, tt.args.left, tt.args.right)
		})
	}
}
