package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_check1752(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1",
			args: args{nums: []int{3, 4, 5, 1, 2}},
			want: true,
		},
		{
			name: "Test2",
			args: args{nums: []int{2, 1, 3, 4}},
			want: false,
		},
		{
			name: "Test3",
			args: args{nums: []int{1, 2, 3}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, check1752(tt.args.nums), "check1752(%v)", tt.args.nums)
		})
	}
}
