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

	merged := ShuffleMerge(a, b)

	fmt.Print("merged: ")
	list.Print(merged)
	fmt.Println()
}

func ShuffleMerge(a, b *list.Node) *list.Node {

	var dummy *list.Node

	dummy = &list.Node{
		Next: a,
	}
	if a == nil {
		dummy.Next = b
	}

	tail := dummy

	for a != nil && b != nil {
		tail.Next = a
		tail = a
		a = a.Next
		tail.Next = b
		tail = b
		b = b.Next
	}

	tail.Next = a

	if b != nil {
		tail.Next = b
	}

	return dummy.Next
}
