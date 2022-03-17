package p396

func maxRotateFunction(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	s := 0
	s0 := 0
	for i, n := range nums {
		s += n
		s0 += i * n
	}
	ans := s0
	l := len(nums)
	for i := 1; i < l; i ++ {
		cur := s0 + s - l * nums[l - i]
		if cur > ans {
			ans = cur
		}
		s0 = cur
	}
	return ans
}
