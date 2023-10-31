package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findArray(t *testing.T) {
	type args struct {
		pref []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test 1",
			args{[]int{5, 2, 0, 3, 1}},
			[]int{5, 7, 2, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findArray(tt.args.pref), "findArray(%v)", tt.args.pref)
		})
	}
}
