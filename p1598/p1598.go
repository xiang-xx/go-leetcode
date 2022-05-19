package p1598

func minOperations(logs []string) int {
	dep := 0
	for _, s := range logs {
		if s == "./" {
			continue
		} else if s == "../" {
			if dep > 0 {
				dep --
			}
		} else {
			dep ++
		}
	}
	return dep
}
