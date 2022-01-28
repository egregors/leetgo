package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test 1",
			args: args{
				points: [][]int{{1, 3}, {-2, 2}},
				k:      1,
			},
			want: [][]int{{-2, 2}},
		},
		{
			name: "test 2",
			args: args{
				points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
				k:      2,
			},
			want: [][]int{{-2, 4}, {3, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, kClosest(tt.args.points, tt.args.k), "kClosest(%v, %v)", tt.args.points, tt.args.k)
		})
	}
}
