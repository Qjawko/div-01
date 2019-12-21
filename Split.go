package student

func Split(str, charset string) []string {
	result := make([]string, countCharset(str, charset)+1)
	lenStr := StringLen([]rune(str))
	lenCharset := StringLen([]rune(charset))

	j := 0
	last := -1 * lenCharset
	for i, ch := range str {
		if ch == rune(charset[0]) {
			if checkMore(str[i:], charset) {
				result[j] = str[last+lenCharset : i]
				j++
				last = i
			}
		}
	}
	result[j] = str[last+lenCharset : lenStr]

	return result
}

func countCharset(str, charset string) int {
	count := 0
	for i, ch := range str {
		if ch == rune(charset[0]) {
			if checkMore(str[i:], charset) {
				count++
			}
		}
	}

	return count
}

func checkMore(str, charset string) bool {
	for i := 0; i < StringLen([]rune(charset)); i++ {
		if str[i] != charset[i] {
			return false
		}
	}

	return true
}
