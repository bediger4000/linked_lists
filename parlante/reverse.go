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

	Reverse(&head)

	fmt.Print("Reversed:    ")
	list.Print(head)
	fmt.Println()
}

// Reverse - iteratively reverse a linked list
func Reverse(head **list.Node) {
	n := *head
	*head = nil

	for n != nil {
		tmp := n.Next

		n.Next = *head
		*head = n

		n = tmp
	}
}
