package student

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i := 1; i < StringArrLen(array); i++ {
		key := array[i] //key hranit v sebe takushiy symbol
		j := i - 1

		for j >= 0 && f(array[j], key) > 0 {
			array[j+1] = array[j]// sravnivaet i zamenyaet   swwap
			j-- // stop when they swapped
		}
		array[j+1] = key // return ouput ex bac   to abc 
	}
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
