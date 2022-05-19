package main

import "testing"

func Test_smallestSubsequence(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{s: "bcabc"},
			want: "abc",
		},
		{
			name: "case2",
			args: args{"cbacdcbc"}, // c, cb, cba, bac, bacd, bacd, acdb
			want: "acdb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestSubsequence(tt.args.s); got != tt.want {
				t.Errorf("smallestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
