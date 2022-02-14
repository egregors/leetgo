package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rob213(t *testing.T) {
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
				nums: []int{2, 3, 2},
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: 4,
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{1, 2, 3, 1, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, rob213(tt.args.nums), "rob213(%v)", tt.args.nums)
		})
	}
}
