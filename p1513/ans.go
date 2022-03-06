package p1513

const (
	prim = 1000000007
)

func numSub(s string) int {
	ans := 0
	last := 0

	for _, c := range s {
		if c == '0' {
			last = 0
		} else {
			last = last + 1
			if last > prim {
				last = last % prim
			}
			ans += last
			if ans > prim {
				ans = ans % prim
			}
		}
	}
	return ans
}

func main() {

}
