package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search33(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "Test 2",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "Test 3",
			args: args{
				nums:   []int{1, 3},
				target: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, search33(tt.args.nums, tt.args.target), "search33(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}
