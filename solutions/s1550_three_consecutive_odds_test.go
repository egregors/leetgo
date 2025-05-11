package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeConsecutiveOdds(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1",
			args: args{
				arr: []int{2, 6, 4, 1},
			},
			want: false,
		},
		{
			name: "Test2",
			args: args{
				arr: []int{1, 2, 34, 3, 4, 5, 7, 23, 12},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, threeConsecutiveOdds(tt.args.arr), "threeConsecutiveOdds(%v)", tt.args.arr)
		})
	}
}
