package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, increasingTriplet(tt.args.nums), "increasingTriplet(%v)", tt.args.nums)
		})
	}
}
