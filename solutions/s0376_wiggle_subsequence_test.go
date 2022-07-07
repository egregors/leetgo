package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wiggleMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "wiggleMaxLength",
			args: args{
				nums: []int{1, 7, 4, 9, 2, 5},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, wiggleMaxLength(tt.args.nums), "wiggleMaxLength(%v)", tt.args.nums)
		})
	}
}
