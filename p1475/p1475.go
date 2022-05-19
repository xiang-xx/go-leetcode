package p1475

func finalPrices(prices []int) []int {
	stack := make([]int, 0)
	for i, p := range prices {
		for len(stack) > 0 && p <= prices[stack[len(stack) - 1]] {
			prices[stack[len(stack) - 1]] -= p
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}
	return prices
}
