package main

import (
	"os"

	student ".."
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if stringArrLen(args) == 0 {
		z01.PrintRune('\n')
	}

	j := 0
	vowel := ""
	str := []rune(combineAllStrings(args))

	for _, ch := range str {
		if isVowel(ch) {
			j++
			vowel += string(ch)
		}
	}

	for i := range str {
		if isVowel(str[i]) {
			j--
			str[i] = rune(vowel[j])
		}
	}

	student.PrintString(string(str) + "\n")
}

func combineAllStrings(strs []string) string {
	result := ""
	for _, s := range strs {
		result += " " + s
	}

	return result[1:]
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'A' || r == 'e' ||
		r == 'E' || r == 'i' || r == 'I' ||
		r == 'o' || r == 'O' || r == 'u' ||
		r == 'U'
}

func stringArrLen(strs []string) int {
	count := 0
	for range strs {
		count++
	}

	return count
}
