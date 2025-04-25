package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countInterestingSubarrays(t *testing.T) {
	type args struct {
		nums   []int
		modulo int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test Case 1",
			args: args{
				nums:   []int{3, 2, 4},
				modulo: 2,
				k:      1,
			},
			want: 3,
		},
		{
			name: "Test Case 2",
			args: args{
				nums:   []int{3, 1, 9, 6},
				modulo: 3,
				k:      0,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countInterestingSubarrays(tt.args.nums, tt.args.modulo, tt.args.k), "countInterestingSubarrays(%v, %v, %v)", tt.args.nums, tt.args.modulo, tt.args.k)
		})
	}
}
