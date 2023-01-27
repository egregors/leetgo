package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findCheapestPrice(t *testing.T) {
	type args struct {
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
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
				flights: [][]int{
					{0, 1, 100},
					{1, 2, 100},
					{2, 0, 100},
					{1, 3, 600},
					{2, 3, 200},
				},
				src: 0,
				dst: 3,
				k:   1,
			},
			want: 700,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findCheapestPrice(tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.k), "findCheapestPrice(%v, %v, %v, %v, %v)", tt.args.n, tt.args.flights, tt.args.src, tt.args.dst, tt.args.k)
		})
	}
}
