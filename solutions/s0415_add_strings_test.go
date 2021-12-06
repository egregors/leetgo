package solutions

import "testing"

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "579",
		},
		{
			name: "Test 2",
			args: args{
				num1: "123",
				num2: "0",
			},
			want: "123",
		},
		{
			name: "Test 3",
			args: args{
				num1: "0",
				num2: "123",
			},
			want: "123",
		},
		{
			name: "Test 4",
			args: args{
				num1: "0",
				num2: "0",
			},
			want: "0",
		},
		{
			name: "Test 5",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "579",
		},
		{
			name: "Test 6",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "579",
		},
		{
			name: "Test 7",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "579",
		},
		{
			name: "Test 8",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "579",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
