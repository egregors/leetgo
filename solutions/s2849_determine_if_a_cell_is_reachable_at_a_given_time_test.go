package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isReachableAtTime(t *testing.T) {
	type args struct {
		sx int
		sy int
		fx int
		fy int
		t  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				sx: 1,
				sy: 1,
				fx: 3,
				fy: 3,
				t:  2,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				sx: 2,
				sy: 4,
				fx: 7,
				fy: 7,
				t:  6,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isReachableAtTime(tt.args.sx, tt.args.sy, tt.args.fx, tt.args.fy, tt.args.t), "isReachableAtTime(%v, %v, %v, %v, %v)", tt.args.sx, tt.args.sy, tt.args.fx, tt.args.fy, tt.args.t)
		})
	}
}
