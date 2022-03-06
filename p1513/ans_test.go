package p1513

import "testing"

func Test_numSub(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{s: "0110111"},
			want: 9,
		},
		{
			name: "case2",
			args: args{s: "101"},
			want: 2,
		},
		{
			name: "case3",
			args: args{"111111"},
			want: 21,
		},
		{
			name: "case4",
			args: args{"000"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSub(tt.args.s); got != tt.want {
				t.Errorf("numSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
