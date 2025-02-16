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
			break
		}
	}

	fmt.Print("a: ")
	list.Print(a)
	fmt.Println()
	fmt.Print("b: ")
	list.Print(b)
	fmt.Println()

	MoveNode(&a, &b)

	fmt.Println("After MoveNode")
	fmt.Print("a: ")
	list.Print(a)
	fmt.Println()
	fmt.Print("b: ")
	list.Print(b)
	fmt.Println()
}

func MoveNode(a, b **list.Node) {
	first := *a
	if *a != nil {
		*a = (*a).Next
	}
	if first != nil {
		first.Next = *b
		*b = first
	}
}
