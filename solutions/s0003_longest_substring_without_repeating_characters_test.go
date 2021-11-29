package solutions

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "Test 3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "Test 4",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "Test 5",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "Test 6",
			args: args{
				s: "au",
			},
			want: 2,
		},
		{
			name: "Test 7",
			args: args{
				s: "abcdefghijklmnopqrstuvwxyz",
			},
			want: 26,
		},
		{
			name: "Test 8",
			args: args{
				s: "qrsvbspk",
			},
			want: 5,
		},
		{
			name: "Test 9",
			args: args{
				s: "aab",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
