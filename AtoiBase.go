package student

func AtoiBase(nbr string, base string) int {
	if checkBase(base) == false {
		return 0
	}

	power := 1 //stepen'
	result := 0
	for i := StringLen([]rune(nbr)) - 1; i >= 0; i-- {
		id := getIdFrom(nbr[i], base)
		if id != -1 {
			result += id * power
			power *= StringLen([]rune(base))
		} else {
			return 0
		}
	}

	return result
}

func getIdFrom(r byte, s string) int {
	for i, ch := range s {
		if ch == rune(r) {
			return i
		}
	}
	return -1
}
