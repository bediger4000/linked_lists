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

	Append(&a, &b)

	fmt.Println()
	fmt.Print("a: ")
	list.Print(a)
	fmt.Println()
	fmt.Print("b: ")
	list.Print(b)
	fmt.Println()
}

func Append(a, b **list.Node) {
	var dummy = &list.Node{
		Next: *a,
	}

	tail := dummy
	for tail != nil && tail.Next != nil {
		tail = tail.Next
	}

	for ; *b != nil; *b = (*b).Next {
		tail.Next = *b
		tail = tail.Next
	}

	*a = dummy.Next
}
