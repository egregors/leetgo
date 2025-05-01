package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeOccurrences(t *testing.T) {
	type args struct {
		s    string
		part string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				s:    "daabcbaabcbc",
				part: "abc",
			},
			want: "dab",
		},
		{
			name: "Test 2",
			args: args{
				s:    "axxxxyyyyb",
				part: "xy",
			},
			want: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeOccurrences(tt.args.s, tt.args.part), "removeOccurrences(%v, %v)", tt.args.s, tt.args.part)
		})
	}
}
