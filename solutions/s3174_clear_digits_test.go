package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_clearDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{s: "abc"},
			want: "abc",
		},
		{
			name: "test2",
			args: args{s: "cb34"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, clearDigits(tt.args.s), "clearDigits(%v)", tt.args.s)
		})
	}
}
