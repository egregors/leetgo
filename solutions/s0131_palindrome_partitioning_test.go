package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_partition131(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"1",
			args{s: "aab"},
			[][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			"2",
			args{s: "a"},
			[][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, partition131(tt.args.s), "partition131(%v)", tt.args.s)
		})
	}
}
