package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumWealth(t *testing.T) {
	type args struct {
		accounts [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				accounts: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: 9,
		},
		{
			name: "Example2",
			args: args{
				accounts: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
					{10, 11, 12},
				},
			},
			want: 12,
		},
		{
			name: "Example3",
			args: args{
				accounts: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
					{10, 11, 12},
					{13, 14, 15},
				},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumWealth(tt.args.accounts), "maximumWealth(%v)", tt.args.accounts)
		})
	}
}
