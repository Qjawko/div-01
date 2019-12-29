package student

import "fmt"

func SplitWhiteSpaces(str string) []string {
	str = CleanWhiteSpaces(str)
	len := StringLen([]rune(str))
	spaces := countChars(str, ' ')

	result := make([]string, spaces+1)
	j := 0
	lastSpace := -1

	for i, ch := range str {
		if ch == ' ' {
			result[j] = str[lastSpace+1 : i]
			j++
			lastSpace = i
		}
	}
	result[j] = str[lastSpace+1 : len]

	return result
}

func CleanWhiteSpaces(str string) string {
	count := 0
	for _, ch := range str {
		if ch == ' ' || ch == '\t' || ch == '\n' {
			count++
		} else {
			break
		}
	}

	i := StringLen([]rune(str)) - 1
	for ; i >= 0; i-- {
		if str[i] != ' ' || str[i] != '\t' || str[i] != '\n' {
			break
		}
	}

	str = str[count : i+1]
	result := ""

	for i := range str {
		if i + 1>= StringLen([]rune(str)) {
			result += string(str[i])
			break
		}
		if checkForWhiteSpaces(str[i], str[i+1]) {
			continue
		} else {
			result += string(str[i])
		}
	}

	return result
}

func countChars(str string, r rune) int {
	count := 0
	for _, ch := range str {
		if ch == r {
			count++
		}
	}

	return count
}

func checkForWhiteSpaces(b, next byte) bool {
	return b == ' ' && next == ' ' ||
		b == '\t' && next == '\t' ||
		b == '\n' && next == '\n' ||
		b == ' ' && next == '\t' ||
		b == ' ' && next == '\n' ||
		b == '\t' && next == ' ' ||
		b == '\t' && next == '\n' ||
		b == '\n' && next == ' ' ||
		b == '\n' && next == '\t'
}
