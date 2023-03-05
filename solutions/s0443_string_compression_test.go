package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, compress(tt.args.chars), "compress(%v)", tt.args.chars)
		})
	}
}
