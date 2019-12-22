package student

func ActiveBits(n int) uint {
	var result uint = 0
	for n != 0 {
		if n%2 != 0 {
			result++
		}
		n /= 2
	}

	return result
}
