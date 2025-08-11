package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productQueries(t *testing.T) {
	type args struct {
		n       int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				n:       15,
				queries: [][]int{{0, 1}, {2, 2}, {0, 3}},
			},
			want: []int{2, 4, 64},
		}, {
			name: "Test 2",
			args: args{
				n:       2,
				queries: [][]int{{0, 0}},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, productQueries(tt.args.n, tt.args.queries), "productQueries(%v, %v)", tt.args.n, tt.args.queries)
		})
	}
}
