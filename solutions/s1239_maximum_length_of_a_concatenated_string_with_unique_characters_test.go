package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxLength(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-1",
			args: args{
				arr: []string{"un", "iq", "ue"},
			},
			want: 4,
		},
		{
			name: "test-2",
			args: args{
				arr: []string{"cha", "r", "act", "ers"},
			},
			want: 6,
		},
		{
			name: "test-3",
			args: args{
				arr: []string{"abcdefghijklmnopqrstuvwxyz"},
			},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxLength(tt.args.arr), "maxLength(%v)", tt.args.arr)
		})
	}
}
