package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildArray(t *testing.T) {
	type args struct {
		target []int
		n      int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example 1",
			args: args{
				target: []int{1, 3},
				n:      3,
			},
			want: []string{"Push", "Push", "Pop", "Push"},
		},
		{
			name: "Example 2",
			args: args{
				target: []int{1, 2, 3},
				n:      3,
			},
			want: []string{"Push", "Push", "Push"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, buildArray(tt.args.target, tt.args.n), "buildArray(%v, %v)", tt.args.target, tt.args.n)
		})
	}
}
