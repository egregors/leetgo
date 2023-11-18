package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxFrequency(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				nums: []int{1, 2, 4},
				k:    5,
			},
			want: 3,
		},
		{
			name: "case 2",
			args: args{
				nums: []int{1, 4, 8, 13},
				k:    5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxFrequency(tt.args.nums, tt.args.k), "maxFrequency(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
