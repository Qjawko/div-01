package main

import (
	"fmt"

	student ".."
)

func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	// var link *student.NodeI
	// var link2 *student.NodeI

	// link = listPushBack(link, 3)
	// link = listPushBack(link, 5)
	// link = listPushBack(link, 7)

	// link2 = listPushBack(link2, -2)
	// link2 = listPushBack(link2, 9)

	// PrintList(student.SortedListMerge(link2, link))

	nbits := student.ActiveBits(55)
	fmt.Println(nbits)
}