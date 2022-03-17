package p393

func validUtf8(data []int) bool {
	lastCount := 0
	for _, b := range data {
		if lastCount > 0 {
			if b & 0b11000000 != 0b10000000 {
				return false
			}
			lastCount --
			continue
		} else {
			if b & 0b10000000 == 0 {
				continue
			} else if b & 0b11100000 == 0b11000000 {
				lastCount = 1
			} else if b & 0b11110000 == 0b11100000 {
				lastCount = 2
			} else if b & 0b11111000 == 0b11110000 {
				lastCount = 3
			} else {
				return false
			}
		}
	}
	if lastCount > 0 {
		return false
	}
	return true
}
