package student

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i := 1; i < StringArrLen(array); i++ {
		key := array[i]
		j := i - 1

		for j >= 0 && f(array[j], key) > 0 {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
}

func SwapStrings(a, b *string) {
	buffer := *a
	*a = *b
	*b = buffer
}

func StringArrLen(array []string) int {
	count := 0
	for range array {
		count++
	}

	return count
}

func Compare(a, b string) int {
	if a > b {
		return 1
	} else if b > a {
		return -1
	}

	return 0
}
