package main

import "testing"

func Test_nearestPalindromic(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{"123"},
			want: "121",
		},
		{
			name: "case2",
			args: args{"1"},
			want: "0",
		},
		{
			name: "case3",
			args: args{"2"},
			want: "1",
		},
		{
			name: "case4",
			args: args{"12345654641"},
			want: "12345654321",
		},
		{
			name: "case5",
			args: args{"10"},
			want: "9",
		},
		{
			name: "case6",
			args: args{"11"},
			want: "9",
		},
		{
			name: "case7",
			args: args{"111"},
			want: "101",
		},
		{
			name: "case8",
			args: args{"120"},
			want: "121",
		},
		{
			name: "case9",
			args: args{"1000"},
			want: "999",
		},
		{
			name: "case9",
			args: args{"9999"},
			want: "10001",
		},
		{
			name: "case10",
			args: args{"10001"},
			want: "9999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestPalindromic(tt.args.n); got != tt.want {
				t.Errorf("nearestPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}
