package p393

import "testing"

func Test_validUtf8(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"1byte true",
			args: args{
				data: []int{0b00000001},
			},
			want: true,
		},
		{
			name: "1byte false",
			args: args{
				data: []int{0b10000001, 0b10000000},
			},
			want: false,
		},
		{
			name: "2byte true",
			args: args{
				data: []int{0b11010101, 0b10000101, 0b00000010},
			},
			want: true,
		},
		{
			name: "2byte false",
			args: args{
				data: []int{0b00000010, 0b11010101, 0b11000101, 0b00000010},
			},
			want: false,
		},
		{
			name: "3byte true",
			args: args{
				data: []int{0b11100101, 0b10000101, 0b10000010, 0b00000010},
			},
			want: true,
		},
		{
			name: "3byte false",
			args: args{
				data: []int{0b00000010, 0b11100101, 0b10000101, 0b00000010},
			},
			want: false,
		},
		{
			name: "4byte true",
			args: args{
				data: []int{0b11110101, 0b10000101, 0b10000010, 0b10000010, 0b00000010},
			},
			want: true,
		},
		{
			name: "4byte false",
			args: args{
				data: []int{0b00000010, 0b11110101, 0b10000101, 0b10000010, 0b00000010},
			},
			want: false,
		},
		{
			name: "mix true",
			args: args{
				data: []int{0b11110101, 0b10000101, 0b10000010, 0b10000010, 0b00000010, 0b11100101, 0b10000101, 0b10000010, 0b00000010},
			},
			want: true,
		},
		{
			name: "mix false",
			args: args{
				data: []int{0b11110101, 0b10000101, 0b10000010, 0b10000010, 0b00000010, 0b11100101, 0b11000101, 0b10000010, 0b00000010},
			},
			want: false,
		},
		{
			name: "mix false",
			args: args{
				data: []int{0b11110101, 0b10000101, 0b10000010, 0b10000010, 0b00000010, 0b11100101, 0b11000101, 0b10000010, 0b11000010},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validUtf8(tt.args.data); got != tt.want {
				t.Errorf("validUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}
