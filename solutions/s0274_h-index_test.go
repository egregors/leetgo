package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				citations: []int{3, 0, 6, 1, 5},
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				citations: []int{1, 3, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, hIndex(tt.args.citations), "hIndex(%v)", tt.args.citations)
		})
	}
}
