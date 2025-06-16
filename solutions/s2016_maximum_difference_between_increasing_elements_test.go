package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{nums: []int{7, 1, 5, 4}},
			want: 4,
		},
		{
			name: "Test2",
			args: args{nums: []int{9, 4, 3, 2}},
			want: -1,
		},
		{
			name: "Test3",
			args: args{nums: []int{1, 5, 2, 10}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumDifference(tt.args.nums), "maximumDifference(%v)", tt.args.nums)
		})
	}
}
