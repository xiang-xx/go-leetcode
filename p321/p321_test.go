package main

import (
	"reflect"
	"testing"
)

func Test_maxNumber(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				nums1: []int{3,4,6,5},
				nums2: []int{9, 1, 2, 5, 8, 3},
				k: 5,
			},
			want: []int{9, 8, 6, 5, 3},
		},
		{
			name: "case2",
			args: args{
				nums1: []int{3,4,6,5},
				nums2: []int{9, 1, 2, 5, 8, 3},
				k: 6,
			},
			want: []int{9, 8, 4, 6, 5, 3},
		},
		{
			name: "case3",
			args: args{
				nums1: []int{3,4,6,5},
				nums2: []int{9, 1, 2, 5, 8, 3},
				k: 2,
			},
			want: []int{9, 8},
		},
		{
			name: "case4",
			args: args{
				nums1: []int{3,4},
				nums2: []int{9, 1},
				k: 4,
			},
			want: []int{9,3,4,1},
		},
		{
			name: "case5",
			args: args{
				nums1: []int{3,4},
				nums2: []int{9, 1},
				k: 3,
			},
			want: []int{9,4,1},
		},
		{
			name: "case6",
			args: args{
				nums1: []int{6,7},
				nums2: []int{6,0,4},
				k: 5,
			},
			want: []int{6,7,6,0,4},
		},
		{
			name: "case7",
			args: args{
				nums1: []int{6,0,4},
				nums2: []int{6,7},
				k: 5,
			},
			want: []int{6,7,6,0,4},
		},
		{
			name: "case8",
			args: args{
				nums1: []int{6,7},
				nums2: []int{6,7,8},
				k: 5,
			},
			want: []int{6,7,8,6,7},
		},
		{
			name: "case9",
			args: args{
				nums1: []int{6,7,8},
				nums2: []int{6,7},
				k: 5,
			},
			want: []int{6,7,8,6,7},
		},
		{
			name: "case10",
			args: args{
				nums1: []int{8,9},
				nums2: []int{3,9},
				k: 3,
			},
			want: []int{9,8,9},
		},
		{
			name: "case11",
			args: args{
				nums1: []int{8,0,4,4,1,7,3,6,5,9,3,6,6,0,2,5,1,7,7,7,8,7,1,4,4,5,4,8,7,6,2,2,9,4,7,5,6,2,2,8,4,6,0,4,7,8,9,1,7,0},
				nums2: []int{6,9,8,1,1,5,7,3,1,3,3,4,9,2,8,0,6,9,3,3,7,8,3,4,2,4,7,4,5,7,7,2,5,6,3,6,7,0,3,5,3,2,8,1,6,6,1,0,8,4},
				k: 50,
			},
			want: []int{9,9,9,9,9,8,7,5,6,3,4,2,4,7,4,5,7,7,2,5,6,3,6,7,2,2,8,4,6,0,4,7,8,9,1,7,0,3,5,3,2,8,1,6,6,1,0,8,4,0},
		},
		{
			name: "case12",
			args: args{
				nums2: []int{8,0,4,4,1,7,3,6,5,9,3,6,6,0,2,5,1,7,7,7,8,7,1,4,4,5,4,8,7,6,2,2,9,4,7,5,6,2,2,8,4,6,0,4,7,8,9,1,7,0},
				nums1: []int{6,9,8,1,1,5,7,3,1,3,3,4,9,2,8,0,6,9,3,3,7,8,3,4,2,4,7,4,5,7,7,2,5,6,3,6,7,0,3,5,3,2,8,1,6,6,1,0,8,4},
				k: 50,
			},
			want: []int{9,9,9,9,9,8,7,5,6,3,4,2,4,7,4,5,7,7,2,5,6,3,6,7,2,2,8,4,6,0,4,7,8,9,1,7,0,3,5,3,2,8,1,6,6,1,0,8,4,0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumber(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
