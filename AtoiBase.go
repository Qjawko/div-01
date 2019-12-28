package student

func AtoiBase(nbr string, base string) int {
	if checkBase(base) == false {
		return 0
	}
	power := 1
	result := 0
	for i := StringLen([]rune(nbr)) - 1; i >= 0; i-- {
		id := getIdFrom(rune(nbr[i]), base)
		if id != -1 {
			result += id * power
			power *= StringLen([]rune(base))
		} else {
			return 0
		}
	}

	return result
}

func getIdFrom(r rune, s string) int {
	for i, ch := range s {
		if ch == r {
			return i
		}
	}
	return -1
}
