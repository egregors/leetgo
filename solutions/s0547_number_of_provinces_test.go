package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findCircleNum(t *testing.T) {
	type args struct {
		isConnected [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "findCircleNum test",
			args: args{
				isConnected: [][]int{
					{1, 1, 0},
					{1, 1, 0},
					{0, 0, 1},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findCircleNum(tt.args.isConnected), "findCircleNum(%v)", tt.args.isConnected)
		})
	}
}
