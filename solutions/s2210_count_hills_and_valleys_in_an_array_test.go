package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countHillValley(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{nums: []int{2, 4, 1, 1, 6, 5}},
			want: 3,
		},
		{
			name: "test2",
			args: args{nums: []int{6, 6, 5, 5, 4, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countHillValley(tt.args.nums), "countHillValley(%v)", tt.args.nums)
		})
	}
}
