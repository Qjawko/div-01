package student

import "github.com/01-edu/z01"

func ConvertBase(nbr, baseFrom, baseTo string) string {
	atoiResult := AtoiBase(nbr, baseFrom)
	return PrintNbrBaseMod(atoiResult, baseTo)
}

func PrintNbrBaseMod(nbr int, base string) string {
	if !checkBase(base) {
		PrintString("NV")
		return "0"
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

	return result
}
