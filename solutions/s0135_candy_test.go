package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				ratings: []int{1, 0, 2},
			},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{
				ratings: []int{1, 2, 2},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, candy(tt.args.ratings), "candy(%v)", tt.args.ratings)
		})
	}
}
