package solutions

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 0",
			args: args{
				s: "abccccdd",
			},
			want: 7,
		},
		{
			name: "Test 1",
			args: args{
				s: "babad",
			},
			want: 5,
		},
		{
			name: "Test 2",
			args: args{
				s: "cbbd",
			},
			want: 3,
		},
		{
			name: "Test 3",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "Test 4",
			args: args{
				s: "ab",
			},
			want: 1,
		},
		{
			name: "Test 5",
			args: args{
				s: "abcd",
			},
			want: 1,
		},
		{
			name: "Test 6",
			args: args{
				s: "aab",
			},
			want: 3,
		},
		{
			name: "Test 7",
			args: args{
				s: "aabb",
			},
			want: 4,
		},
		{
			name: "Test 8",
			args: args{
				s: "aabbcc",
			},
			want: 6,
		},
		{
			name: "Test 9",
			args: args{
				s: "aabbccdd",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
