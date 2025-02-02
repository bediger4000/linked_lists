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

	intersection := SortedIntersect(a, b)

	fmt.Print("intersection: ")
	list.Print(intersection)
	fmt.Println()
}

func SortedIntersect(a, b *list.Node) *list.Node {

	if a == nil || b == nil {
		return nil
	}

	first, second := a, b
	if b.Data < a.Data {
		first, second = b, a
	}

	for ; first.Data < second.Data; first = first.Next {
	}

}
