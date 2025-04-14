package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countGoodTriplets(t *testing.T) {
	type args struct {
		arr []int
		a   int
		b   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				arr: []int{3, 0, 1, 1, 9, 7},
				a:   7,
				b:   2,
				c:   3,
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				arr: []int{1, 1, 2, 2, 3},
				a:   0,
				b:   0,
				c:   1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countGoodTriplets(tt.args.arr, tt.args.a, tt.args.b, tt.args.c), "countGoodTriplets(%v, %v, %v, %v)", tt.args.arr, tt.args.a, tt.args.b, tt.args.c)
		})
	}
}
