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

	for second != nil {
		// swap two nodes
		first.Next = second.Next
		second.Next = first

		// keep the list in order.
		// need a node to point to second
		// that currently points to first.
		*previous = second
		// get ready for the next pass through the
		// loop. Might need to update first.Next with
		// the address of a swapped node.
		previous = &(first.Next)

		// Advance first and second pointers.
		// first.Next holds the next node in the list.
		first = first.Next
		second = nil
		if first != nil {
			// since first has advanced, set second appropriately.
			second = first.Next
		}
	}

	return head
}
