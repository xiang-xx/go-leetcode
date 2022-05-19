package main

import "strings"

func smallestSubsequence(s string) string {
	stack := make([]int32, 0)
Outer:
	for i, c := range s {
		// 已经在 stack 的忽略
		for _, tmp := range stack {
			if tmp == c {
				continue Outer
			}
		}
		// 当前比栈顶小，进入循环
		for len(stack) > 0 && c < stack[len(stack) - 1] {
			find := false
			// 后续元素有栈顶元素，pop 栈顶，继续判断新的栈顶
			for j := i + 1; j < len(s); j ++ {
				if int32(s[j]) == stack[len(stack) - 1] {
					stack = stack[:len(stack) - 1]
					find = true
					break
				}
			}
			// 后续没有栈顶元素，跳出循环
			if !find {
				break
			}
		}
		stack = append(stack, c)
	}
	sb := strings.Builder{}
	for _, c := range stack {
		sb.WriteRune(c)
	}
	return sb.String()
}

func main() {
	println(smallestSubsequence("cbacdcbc"))
}