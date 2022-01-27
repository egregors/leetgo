package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findJudge(t *testing.T) {
	type args struct {
		n     int
		trust [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				n: 2,
				trust: [][]int{
					{1, 2},
				},
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				n: 3,
				trust: [][]int{
					{1, 3},
					{2, 3},
				},
			},
			want: 3,
		},
		{
			name: "Test 3",
			args: args{
				n: 3,
				trust: [][]int{
					{1, 3},
					{2, 3},
					{3, 1},
				},
			},
			want: -1,
		},
		{
			name: "Test 4",
			args: args{
				n: 3,
				trust: [][]int{
					{1, 2},
					{2, 3},
				},
			},
			want: -1,
		},
		{
			name: "Test 5",
			args: args{
				n: 4,
				trust: [][]int{
					{1, 3},
					{1, 4},
					{2, 3},
					{2, 4},
					{4, 3},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findJudge(tt.args.n, tt.args.trust), "findJudge(%v, %v)", tt.args.n, tt.args.trust)
		})
	}
}
