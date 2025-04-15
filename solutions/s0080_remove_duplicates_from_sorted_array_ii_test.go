package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeDuplicates80(t *testing.T) {
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
			args: args{nums: []int{1, 1, 1, 2, 2, 3}},
			want: 5,
		},
		{
			name: "test2",
			args: args{nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeDuplicates80(tt.args.nums), "removeDuplicates80(%v)", tt.args.nums)
		})
	}
}
