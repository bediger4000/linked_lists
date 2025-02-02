package main

import (
	"fmt"
	"linked_lists/list"
	"os"
)

func main() {

	head := list.Compose(os.Args[1:])
	fmt.Print("source list: ")
	list.Print(head)
	fmt.Println("\n")

	hd, _ := RecursiveReverse(head)

	fmt.Print("Reversed:    ")
	list.Print(hd)
	fmt.Println()
}

func RecursiveReverse(node *list.Node) (*list.Node, *list.Node) {
	if node.Next == nil {
		return node, node // head, tail of reversed list
	}

	head, tail := RecursiveReverse(node.Next)
	tail.Next = node
	tail = node
	tail.Next = nil

	return head, tail
}
