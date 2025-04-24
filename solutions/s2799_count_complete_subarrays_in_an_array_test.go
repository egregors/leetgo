package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countCompleteSubarrays(t *testing.T) {
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
			args: args{nums: []int{1, 3, 1, 2, 2}},
			want: 4,
		},
		{
			name: "Test 2",
			args: args{nums: []int{5, 5, 5, 5}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countCompleteSubarrays(tt.args.nums), "countCompleteSubarrays(%v)", tt.args.nums)
		})
	}
}
