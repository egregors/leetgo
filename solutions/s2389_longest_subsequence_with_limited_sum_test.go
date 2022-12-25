package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_answerQueries(t *testing.T) {
	type args struct {
		nums    []int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				nums:    []int{4, 5, 2, 1},
				queries: []int{3, 10, 21},
			},
			want: []int{2, 3, 4},
		},
		{
			name: "test 2",
			args: args{
				nums:    []int{2, 3, 4, 5},
				queries: []int{1},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, answerQueries(tt.args.nums, tt.args.queries), "answerQueries(%v, %v)", tt.args.nums, tt.args.queries)
		})
	}
}
