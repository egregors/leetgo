package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calPoints(t *testing.T) {
	type args struct {
		ops []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				ops: []string{"5", "2", "C", "D", "+"},
			},
			want: 30,
		},
		{
			name: "Test 2",
			args: args{
				ops: []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, calPoints(tt.args.ops), "calPoints(%v)", tt.args.ops)
		})
	}
}
