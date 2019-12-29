package student

func PrintCombN(n int) {

	if n < 1 || n > 9 {
		PrintString("\n")
		return
	}

	current := getFirstNumber(n)

	lastNumber := getLastNumber(n)

	for current != lastNumber {
		PrintString(current + ", ")
		current = getNextNumber(current)
	}

	PrintString(current + "\n")
}

func getNextNumber(curr string) string {

	current := []rune(curr)

	for {
		current = increment(current)
		if isSorted(current) {
			break
		}
	}

	return string(current)
}

func increment(s []rune) []rune {
	len := 0
	for range s {
		len++
	}

	if s[len-1] != '9' {
		s[len-1]++
	} else {
		f := increment(s[:len-1])
		return []rune(string(f) + "0")
	}

	return s
}

func isSorted(s []rune) bool {
	for i := range s {
		if i != 0 && s[i-1] >= s[i] {
			return false
		}
	}

	return true
}

func getFirstNumber(n int) string {
	r := ""

	for i := 0; i < n; i++ {
		r += string(rune(i + '0'))
	}

	return r
}

func getLastNumber(n int) string {
	r := ""
	for i := 0; i < n; i++ {
		r = string(rune(9-i+'0')) + r
	}
	return r
}
