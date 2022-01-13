package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, productExceptSelf(tt.args.nums), "productExceptSelf(%v)", tt.args.nums)
		})
	}
}
