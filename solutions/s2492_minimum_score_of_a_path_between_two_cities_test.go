package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minScore(t *testing.T) {
	type args struct {
		n     int
		roads [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				4,
				[][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minScore(tt.args.n, tt.args.roads), "minScore(%v, %v)", tt.args.n, tt.args.roads)
		})
	}
}
