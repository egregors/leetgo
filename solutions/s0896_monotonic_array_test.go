package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isMonotonic(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{1, 2, 2, 3},
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{6, 5, 4, 4},
			},
			want: true,
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{1, 3, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isMonotonic(tt.args.nums), "isMonotonic(%v)", tt.args.nums)
		})
	}
}
