package student

func Atoi(s string) int {
	result := 0

	str := []rune(s)
	isNegative := false
	if s[0] == '-' || s[0] == '+' {
		str = []rune(s[1:])
		if s[0] == '-' {
			isNegative = true
		}
	}
	len := stringLen(str)

	power := 1

	for i := len - 1; i >= 0; i-- {
		//z01.PrintRune(str[i])
		if str[i] >= '0' && str[i] <= '9' {
			result += int(str[i]-48) * power
			power *= 10
		} else {
			return 0
		}
	}

	if isNegative {
		return result * -1
	}

	return result
}

func stringLen(s []rune) int {
	count := 0
	for range s {
		count++
	}

	return count
}
