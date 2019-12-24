package main

import (
	"fmt"

	student ".."
)

func PrintList(l *student.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
	link := &student.List{}
	link2 := &student.List{}

	fmt.Println("----normal state----")
	student.ListPushBack(link2, 1)
	PrintList(link2)
	student.ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "Hello")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "There")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "How")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "are")
	student.ListPushBack(link, "you")
	student.ListPushBack(link, 1)
	PrintList(link)

	student.ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}
