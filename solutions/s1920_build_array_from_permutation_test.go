package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_buildArray0005(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{0, 2, 1, 5, 3, 4},
			},
			want: []int{0, 1, 2, 4, 5, 3},
		},
		{
			name: "test2",
			args: args{
				nums: []int{5, 0, 1, 2, 3, 4},
			},
			want: []int{4, 5, 0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, buildArray0005(tt.args.nums), "buildArray0005(%v)", tt.args.nums)
		})
	}
}
