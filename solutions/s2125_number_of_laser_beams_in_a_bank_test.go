package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfBeams(t *testing.T) {
	type args struct {
		bank []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{bank: []string{"011001", "000000", "010100", "001000"}},
			want: 8,
		},
		{
			name: "test2",
			args: args{bank: []string{"000", "111", "000"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numberOfBeams(tt.args.bank), "numberOfBeams(%v)", tt.args.bank)
		})
	}
}
