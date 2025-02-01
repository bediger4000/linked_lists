package main

import (
	"fmt"
	"linked_lists/list"
	"os"
)

func main() {

	var a, b *list.Node

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "--" {
			a = list.Compose(os.Args[1:i])
			b = list.Compose(os.Args[i:])
		}
	}

	fmt.Print("a: ")
	list.Print(a)
	fmt.Println()
	fmt.Print("b: ")
	list.Print(b)
	fmt.Println()

	merged := SortedMerge(a, b)

	fmt.Print("merged: ")
	list.Print(merged)
	fmt.Println()
}

func SortedMerge(a, b *list.Node) *list.Node {

	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	x := &a
	if b.Data < a.Data {
		x = &b
	}

	head, tail := *x, *x

	*x = (*x).Next

	for a != nil && b != nil {
		n := &a
		if b.Data < a.Data {
			n = &b
		}
		tail.Next = *n
		*n = (*n).Next
		tail = tail.Next
	}

	tail.Next = a

	if b != nil {
		tail.Next = b
	}

	return head
}
