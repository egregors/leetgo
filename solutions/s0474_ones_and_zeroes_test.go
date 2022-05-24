package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMaxForm(t *testing.T) {
	type args struct {
		strs []string
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				strs: []string{"10", "0001", "111001", "1", "0"},
				m:    5,
				n:    3,
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				strs: []string{"10", "0", "1"},
				m:    1,
				n:    1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMaxForm(tt.args.strs, tt.args.m, tt.args.n), "findMaxForm(%v, %v, %v)", tt.args.strs, tt.args.m, tt.args.n)
		})
	}
}
