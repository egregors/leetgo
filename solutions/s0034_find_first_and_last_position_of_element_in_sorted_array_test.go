package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				nums:   []int{0},
				target: 0,
			},
			want: []int{0, 0},
		},
		{
			name: "Test 2",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: []int{-1, -1},
		},
		{
			name: "Test 3",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 10,
			},
			want: []int{5, 5},
		},
		{
			name: "Test 4",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 5,
			},
			want: []int{0, 0},
		},
		{
			name: "Test 5",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 7,
			},
			want: []int{1, 2},
		},
		{
			name: "Test 6",
			args: args{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, searchRange(tt.args.nums, tt.args.target), "searchRange(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}
