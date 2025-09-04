package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findClosest(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{x: 2, y: 7, z: 4},
			want: 1,
		},
		{
			name: "test2",
			args: args{x: 2, y: 5, z: 6},
			want: 2,
		},
		{
			name: "test3",
			args: args{x: 1, y: 5, z: 3},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findClosest(tt.args.x, tt.args.y, tt.args.z), "findClosest(%v, %v, %v)", tt.args.x, tt.args.y, tt.args.z)
		})
	}
}
