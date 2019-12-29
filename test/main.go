package main

import (
	"fmt"

	student ".."
)

func main() {
	str := "Hello	\n how are you"
	fmt.Println(student.SplitWhiteSpaces(str))
}
