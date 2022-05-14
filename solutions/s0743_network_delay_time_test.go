package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_networkDelayTime(t *testing.T) {
	type args struct {
		times [][]int
		n     int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				times: [][]int{
					{2, 1, 1},
					{2, 3, 1},
					{3, 4, 1},
				},
				n: 4,
				k: 2,
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				times: [][]int{
					{1, 2, 1},
				},
				n: 2,
				k: 1,
			},
			want: 1,
		},
		{
			name: "Test 3",
			args: args{
				times: [][]int{
					{1, 2, 1},
				},
				n: 2,
				k: 2,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, networkDelayTime(tt.args.times, tt.args.n, tt.args.k), "networkDelayTime(%v, %v, %v)", tt.args.times, tt.args.n, tt.args.k)
		})
	}
}
