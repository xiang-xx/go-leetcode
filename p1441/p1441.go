package p1441

func buildArray(target []int, n int) []string {
	index := 0
	ans := make([]string, 0)
	for i := 1; i <= n; i ++ {
		if index >= len(target) {
			break
		}
		if target[index] == i {
			ans = append(ans, "Push")
			index ++
		} else {
			ans = append(ans, "Push", "Pop")
		}
	}
	return ans
}
