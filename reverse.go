package main

import (
	"fmt"
	"os"

	"linked_lists/list"
)

func main() {

	var list1 *list.Node

	list1 = list.Compose(os.Args[1:])

	list.Print(list1)
	fmt.Println()

	reversed := Reverse(list1)
	list.Print(reversed)
	fmt.Println()
}

func Reverse(head *list.Node) (reversed *list.Node) {
	var tmp *list.Node
	for head != nil {
		tmp, head = head, head.Next
		tmp.Next, reversed = reversed, tmp
	}
	return
}
