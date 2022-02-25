package p1

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{nums: []int{2,7,11,15}, target: 9},
			want: []int{0,1},
		},
		{
			args: args{nums: []int{9,3,4,6,10}, target: 7},
			want: []int{1,2},
		},
		{
			args: args{nums: []int{1,2,2,5,6}, target: 4},
			want: []int{1,2},
		},
		{
			args: args{nums: []int{2,7}, target: 9},
			want: []int{0,1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
