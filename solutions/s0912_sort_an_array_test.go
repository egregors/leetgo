package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test0",
			args: args{
				nums: []int{10, 80, 30, 90, 40, 50, 70},
			},
			want: []int{10, 30, 40, 50, 70, 80, 90},
		},
		{
			name: "test1",
			args: args{
				nums: []int{5, 2, 3, 1},
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "test2",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			want: []int{0, 0, 1, 1, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sortArray(tt.args.nums), "sortArray(%v)", tt.args.nums)
		})
	}
}
