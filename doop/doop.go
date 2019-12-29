package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"

	student ".."
)

const INT_MAX = 9223372036854775807
const INT_MIN = -9223372036854775807

func main() {
	args := os.Args[1:]

	if StringArrLen(args) != 3 {
		return
	}

	firstNumber, firstIsValid := AtoiV2(args[0])
	secondNumber, secondIsValid := AtoiV2(args[2])

	if !(firstIsValid && secondIsValid) {

		fmt.Println("here?")
		student.PrintString("0\n")
		return
	}

	if checkNumberForOverflow(firstNumber) ||
		checkNumberForOverflow(secondNumber) {
		return
	}

	switch args[1] {
	case "+":

		if firstNumber > 0 && secondNumber > 0 {
			if firstNumber > INT_MAX-secondNumber {
				result := firstNumber + secondNumber
				student.PrintString(Itoa(result))
				z01.PrintRune('\n')
			}
		}

		if firstNumber < 0 && secondNumber < 0 {
			if firstNumber < INT_MIN-secondNumber {
				result := firstNumber + secondNumber
				student.PrintString(Itoa(result))
				z01.PrintRune('\n')
			}
		}

	case "-":
		if firstNumber < 0 && secondNumber > 0 {
			if firstNumber < INT_MIN+secondNumber {
				student.PrintString("0\n")
				return
			}
		}

		if firstNumber > 0 && secondNumber < 0 {
			if firstNumber < INT_MAX+secondNumber {
				student.PrintString("0\n")
				return
			}
		}

		result := firstNumber - secondNumber
		student.PrintString(Itoa(result))
		z01.PrintRune('\n')
	case "*":
		if firstNumber < 0 && secondNumber < 0 {
			if firstNumber < INT_MAX/secondNumber {
				student.PrintString("0\n")
				return
			}
		}
		if firstNumber > 0 && secondNumber > 0 {
			if firstNumber > INT_MAX/secondNumber {
				student.PrintString("0\n")
				return
			}
		}
		if firstNumber < 0 && secondNumber > 0 {
			if firstNumber < INT_MIN/secondNumber {
				student.PrintString("0\n")
				return
			}
		}
		if firstNumber > 0 && secondNumber < 0 {
			if firstNumber > INT_MIN/secondNumber {
				student.PrintString("0\n")
				return
			}
		}

		result := firstNumber * secondNumber
		student.PrintString(Itoa(result))
		z01.PrintRune('\n')
	case "/":
		if secondNumber == 0 {
			student.PrintString("No division by 0\n")
			return
		}
		result := firstNumber / secondNumber
		student.PrintString(Itoa(result))
		z01.PrintRune('\n')
	case "%":
		if secondNumber == 0 {
			student.PrintString("No modulo by 0\n")
			return
		}
		result := firstNumber % secondNumber
		student.PrintString(Itoa(result))
		z01.PrintRune('\n')
	default:
		student.PrintString("0\n")
	}

}

func checkNumberForOverflow(n int) bool {

	if int64(n) >= 9223372036854775807 || int64(n) <= -9223372036854775807 {
		student.PrintString("0\n")
		return true
	}
	return false
}

func AtoiV2(s string) (int, bool) {
	result := 0

	str := []rune(s)
	isNegative := false
	if s[0] == '-' || s[0] == '+' {
		str = []rune(s[1:])
		if s[0] == '-' {
			isNegative = true
		}
	}
	len := student.StringLen(str)

	power := 1

	for i := len - 1; i >= 0; i-- {
		//z01.PrintRune(str[i])
		if str[i] >= '0' && str[i] <= '9' {
			result += int(str[i]-48) * power
			power *= 10
		} else {
			return 0, false
		}
	}

	if isNegative {
		return result * -1, true
	}

	return result, true
}

func StringArrLen(arr []string) int {
	count := 0
	for range arr {
		count++
	}

	return count
}

func Itoa(nbr int) string {
	result := ""

	isNegative := false
	if nbr < 0 {
		nbr *= -1
		isNegative = true
	}

	for nbr != 0 {
		result = string(nbr%10+48) + result
		nbr /= 10
	}

	if isNegative {
		result = "-" + result
	}

	return result
}
