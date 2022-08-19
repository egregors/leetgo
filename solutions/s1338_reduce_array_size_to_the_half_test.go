package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minSetSize(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				arr: []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				arr: []int{7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minSetSize(tt.args.arr), "minSetSize(%v)", tt.args.arr)
		})
	}
}
