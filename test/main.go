package main

import (
	"fmt"

	student ".."
)

func main() {
	str := "    Hello     how are you?    "
	fmt.Println(student.SplitWhiteSpaces(str))
	fmt.Println(len(student.SplitWhiteSpaces(str)))
}

//111101

//1010111
