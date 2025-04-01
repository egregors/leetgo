package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mostPoints(t *testing.T) {
	type args struct {
		questions [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"case 1",
			args{
				[][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}},
			},
			5,
		},
		{
			"case 2",
			args{
				[][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}},
			},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, mostPoints(tt.args.questions), "mostPoints(%v)", tt.args.questions)
		})
	}
}
