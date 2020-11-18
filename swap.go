package main

import (
	"fmt"
	"linked_lists/list"
	"os"
)

/*
Given the head of a singly linked list,
swap every two nodes and return its head.

For example,
given `1 -> 2 -> 3 -> 4`, return `2 -> 1 -> 4 -> 3`.
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "swap every two nodes of linked list\n")
		os.Exit(1)
	}
	head := list.Compose(os.Args[1:])
	list.Print(head)
	fmt.Println()

	head = swapnodes(head)
	list.Print(head)
	fmt.Println()
}

func swapnodes(head *list.Node) *list.Node {

	if head == nil {
		return nil
	}

	first := head
	second := head.Next
	previous := &head

	if second != nil {
		head = second
	}

	for second != nil {
		first.Next = second.Next
		second.Next = first

		*previous = second
		previous = &(first.Next)

		first = first.Next
		second = nil
		if first != nil {
			second = first.Next
		}
	}

	return head
}
