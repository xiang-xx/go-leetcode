package main

import "math"

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	total := len(nums1) + len(nums2)
	needRemove := total - k
	ans := make([]int, 0)
	index1 := 0
	index2 := 0

	for needRemove > 0 && len(ans) < k {
		// index - index+needRemove 范围内最大元素
		max1, maxIndex1 := maxInRange(nums1, index1, index1 + needRemove)
		max2, maxIndex2 := maxInRange(nums2, index2, index2 + needRemove)
		if max1 > max2 {
			ans = append(ans, max1)
			needRemove = needRemove - (maxIndex1 - index1)
			index1 = maxIndex1 + 1
		} else if max1 < max2 {
			ans = append(ans, max2)
			needRemove = needRemove - (maxIndex2 - index2)
			index2 = maxIndex2 + 1
		} else {
			//if bigger(nums1, index1, nums2, index2) {
			//	ans = append(ans, max2)
			//	needRemove = needRemove - (maxIndex2 - index2)
			//	index2 = maxIndex2 + 1
			//} else {
			//	ans = append(ans, max1)
			//	needRemove = needRemove - (maxIndex1 - index1)
			//	index1 = maxIndex1 + 1
			//}
		}
	}
	if len(ans) == k {
		return ans
	}

	// 余下的数据取取最大值填充足够
	for len(ans) < k {
		if bigger(nums1, index1, nums2, index2) {
			ans = append(ans, nums1[index1])
			index1 ++
		} else {
			ans = append(ans, nums2[index2])
			index2 ++
		}
	}
	return ans
}

func bigger(nums1 []int, index1 int, nums2[]int, index2 int) bool {
	if len(nums1) <= index1 {
		return false
	}
	if len(nums2) <= index2 {
		return true
	}
	if nums1[index1] > nums2[index2] {
		return true
	} else if nums1[index1] < nums2[index2] {
		return false
	} else {
		return bigger(nums1, index1 + 1, nums2, index2 + 1)
	}
}

func maxInRange(nums []int, from, to int) (int, int) {
	maxN := math.MinInt
	maxIndex := from
	for i := from; i <= to && i < len(nums); i ++ {
		if maxN < nums[i] {
			maxN = nums[i]
			maxIndex = i
		}
	}
	return maxN, maxIndex
}

func main() {
	maxNumber([]int{
		8,0,4,4,1,7,3,6,5,9,3,6,6,0,2,5,1,7,7,7,8,7,1,4,4,5,4,8,7,6,2,2,9,4,7,5,6,2,2,8,4,6,0,4,7,8,9,1,7,0},
		[]int{6,9,8,1,1,5,7,3,1,3,3,4,9,2,8,0,6,9,3,3,7,8,3,4,2,4,7,4,5,7,7,2,5,6,3,6,7,0,3,5,3,2,8,1,6,6,1,0,8,4},
		50)
}