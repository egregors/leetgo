package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_destCity(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				paths: [][]string{
					{"London", "New York"},
					{"New York", "Lima"},
					{"Lima", "Sao Paulo"},
				}},
			want: "Sao Paulo",
		},
		{
			name: "Test 2",
			args: args{
				paths: [][]string{
					{"B", "C"},
					{"D", "B"},
					{"C", "A"},
				}},
			want: "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, destCity(tt.args.paths), "destCity(%v)", tt.args.paths)
		})
	}
}
