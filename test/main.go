package main

import (
	student ".."
	"github.com/01-edu/z01"
)

func main() {
	student.PrintNbrBase(125, "876")
	z01.PrintRune('\n')
	student.PrintNbrBase(87, "01")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}

//111101

//1010111
