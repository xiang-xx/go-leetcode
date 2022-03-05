package main

import (
	"math"
	"strconv"
)

// https://leetcode-cn.com/problems/find-the-closest-palindrome/
// 给定一个表示整数的字符串 n ，返回与它最近的回文整数（不包括自身）。如果不止一个，返回较小的那个。
//“最近的”定义为两个整数差的绝对值最小。
//
//示例 1:
//
//输入: n = "123"
//输出: "121"
//示例 2:
//
//输入: n = "1"
//输出: "0"
//解释: 0 和 2是最近的回文，但我们返回最小的，也就是 0。
//
//提示:
//1 <= n.length <= 18
//n只由数字组成
//n不含前导 0
//n代表在[1, 1018- 1] 范围内的整数
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/find-the-closest-palindrome
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// nearestPalindromic 寻找最近的回文数
func nearestPalindromic(n string) string {
	l := len(n)
	candidates := []int{int(math.Pow10(l) + 1), int(math.Pow10(l-1) - 1)}
	half, _ := strconv.Atoi(n[:(l+1)/2])
	for _, h := range []int{half, half + 1, half - 1} {
		y := h
		if l & 1 == 1 {  // l % 2 == 0
			h /= 10
		}
		for h > 0 {
			y = y * 10 + h % 10
			h /= 10
		}
		candidates = append(candidates, y)
	}

	ans := -1
	m := math.MaxInt
	number, _ := strconv.Atoi(n)
	for _, candidate := range candidates {
		if candidate == number {
			continue
		}
		if ans == -1 ||
			abs(candidate - number) < m ||
			abs(candidate - number) == m && candidate < ans {
			ans = candidate
			m = abs(candidate - number)
		}
	}
	return strconv.Itoa(ans)
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func Abs(n int) int {
	x := n >> 63
	return (n ^ x) - x
}

func main() {
	nearestPalindromic("123")
}