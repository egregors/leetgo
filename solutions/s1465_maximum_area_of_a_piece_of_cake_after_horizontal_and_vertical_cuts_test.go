package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea1465(t *testing.T) {
	type args struct {
		h              int
		w              int
		horizontalCuts []int
		verticalCuts   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				h:              5,
				w:              4,
				horizontalCuts: []int{1, 2, 4},
				verticalCuts:   []int{1, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxArea1465(tt.args.h, tt.args.w, tt.args.horizontalCuts, tt.args.verticalCuts), "maxArea1465(%v, %v, %v, %v)", tt.args.h, tt.args.w, tt.args.horizontalCuts, tt.args.verticalCuts)
		})
	}
}
