package p1700

func countStudents(students []int, sandwiches []int) int {
	smap := make(map[int]int)
	for _, s := range students {
		smap[s] ++
	}
	for _, s := range sandwiches {
		if smap[s] > 0 {
			smap[s] --
		} else {
			return smap[0] + smap[1]
		}
	}
	return 0
}
