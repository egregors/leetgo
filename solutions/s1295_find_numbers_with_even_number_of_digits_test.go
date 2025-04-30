package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_findNumbers_1",
			args: args{
				nums: []int{12, 345, 2, 6, 7896},
			},
			want: 2,
		},
		{
			name: "Test_findNumbers_2",
			args: args{
				nums: []int{555, 901, 482, 1771},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findNumbers(tt.args.nums), "findNumbers(%v)", tt.args.nums)
		})
	}
}
