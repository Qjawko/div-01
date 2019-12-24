package main

import (
	"fmt"

	student ".."
)

func main() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	student.BTreeApplyByLevel(root, fmt.Println)
}
