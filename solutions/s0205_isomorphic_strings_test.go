package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{"egg", "add"}, true},
		{"Case 2", args{"foo", "bar"}, false},
		{"Case 3", args{"paper", "title"}, true},
		{"Case 4", args{"ab", "aa"}, false},
		{"Case 5", args{"ab", "ca"}, true},
		{"Case 6", args{"ab", "ab"}, true},
		{"Case 7", args{"ab", "cd"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isIsomorphic(tt.args.s, tt.args.t), "isIsomorphic(%v, %v)", tt.args.s, tt.args.t)
		})
	}
}
