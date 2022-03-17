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
				data: []int{50},
			},
			want: true,
		},
		{
			name: "1byte false",
			args: args{
				data: []int{50, 1<<7, 50},
			},
			want: false,
		},
		{
			name: "2byte true",
			args: args{
				data: []int{0b11},
			},
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
