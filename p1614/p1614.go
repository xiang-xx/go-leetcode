package p1614

func maxDepth(s string) int {
	stack := make([]int32, 0)
	answer := 0
	for _, c := range s {
		if c == '(' {
			stack = append(stack, c)
			answer = maxInt(len(stack), answer)
		} else if c == ')' {
			stack = stack[:len(stack) - 1]
		}
	}
	return answer
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}