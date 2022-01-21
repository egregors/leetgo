package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minRemoveToMakeValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				s: "lee(t(c)o)de)",
			},
			want: "lee(t(c)o)de",
		},
		{
			name: "Test 2",
			args: args{
				s: "a)b(c)d",
			},
			want: "ab(c)d",
		},
		{
			name: "Test 3",
			args: args{
				s: "))(",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minRemoveToMakeValid(tt.args.s), "minRemoveToMakeValid(%v)", tt.args.s)
		})
	}
}
