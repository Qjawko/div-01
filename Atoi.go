package student

import (
	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	len := stringLen(s)

	result := 0

	str := []rune(s)
	isNegative := false
	if s[0] == '-' {
		str = []rune(s[1:])
		isNegative = true
	}

	power := 1

	for i := len - 1; i >= 0; i-- {
		z01.PrintRune(str[i])
		if str[i] >= '0' && str[i] <= '9' {
			result += int(str[i]-48) * power
			power *= 10
		}
	}

	if isNegative {
		return result * -1
	}

	return result
}

func stringLen(s string) int {
	count := 0
	for range s {
		count++
	}

	return count
}
