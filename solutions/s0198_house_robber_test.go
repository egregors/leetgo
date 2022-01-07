package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{2, 7, 9, 3, 1},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, rob(tt.args.nums), "rob(%v)", tt.args.nums)
		})
	}
}
