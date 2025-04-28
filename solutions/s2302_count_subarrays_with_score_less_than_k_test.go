package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countSubarrays2302(t *testing.T) {
	type args struct {
		nums []int
		k    int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{2, 1, 4, 3, 5},
				k:    10,
			},
			want: 6,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{1, 1, 1},
				k:    5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countSubarrays2302(tt.args.nums, tt.args.k), "countSubarrays2302(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
