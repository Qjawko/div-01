package student

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if !checkBase(base) {
		PrintString("NV")
		return
	}

}

func PrintString(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func checkBase(base string) bool {
	if stringLen([]rune(base)) < 2 {
		return false
	}

	return UniqueChars(base)
}

func UniqueChars(s string) bool {
	runes := []rune(s)
	for i, ch := range runes {
		for j := i + 1; j < stringLen(runes); j++ {
			if ch == runes[j] && i != j {
				return false
			}
		}
	}

	return true
}
