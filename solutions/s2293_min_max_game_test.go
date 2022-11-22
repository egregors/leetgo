package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minMaxGame(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{[]int{1, 3, 5, 2, 4, 8, 2, 2}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minMaxGame(tt.args.nums), "minMaxGame(%v)", tt.args.nums)
		})
	}
}
