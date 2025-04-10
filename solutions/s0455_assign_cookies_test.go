package solutions

import "testing"

func Test_findContentChildren(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		g    []int
		s    []int
		want int
	}{
		{
			name: "Test 1",
			g:    []int{1, 2, 3},
			s:    []int{1, 1},
			want: 1,
		},
		{
			name: "Test 2",
			g:    []int{1, 2},
			s:    []int{1, 2, 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findContentChildren(tt.g, tt.s)
			if got != tt.want {
				t.Errorf("findContentChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}
