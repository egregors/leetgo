package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kthCharacter(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		//k=5, b
		//k=10, c
		{
			name: "Test 1",
			args: args{k: 5},
			want: 'b',
		}, {
			name: "Test 2",
			args: args{k: 10},
			want: 'c',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, kthCharacter(tt.args.k), "kthCharacter(%v)", tt.args.k)
		})
	}
}
