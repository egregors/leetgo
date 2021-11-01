package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solve(t *testing.T) {
	input := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'}}
	want := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'}}

	solve(input)
	assert.Equal(t, want, input)

	input = [][]byte{{'X'}}
	want = [][]byte{{'X'}}
	solve(input)
	assert.Equal(t, want, input)
}
