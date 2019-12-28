package student

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if !checkBase(base) {
		PrintString("NV")
		return
	}

	if nbr < 0 {
		nbr *= -1
		z01.PrintRune('-')
	}

	baseLen := StringLen([]rune(base))
	result := ""

	//fmt.Println(baseLen)

	for nbr != 0 {
		result = string(base[nbr%baseLen]) + result
		nbr /= baseLen
	}

	PrintString(result)
}

func PrintString(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func Countains(r rune, s string) bool {
	for _, ch := range s {
		if ch == r {
			return true
		}
	}

	return false
}

func checkBase(base string) bool {
	if StringLen([]rune(base)) < 2 {
		return false
	}

	return UniqueChars(base) && !(Countains('-', base) || Countains('+', base))
}

func UniqueChars(s string) bool {
	runes := []rune(s)
	for i, ch := range runes {
		for j := i + 1; j < StringLen(runes); j++ {
			if ch == runes[j] && i != j {
				return false
			}
		}
	}

	return true
}
