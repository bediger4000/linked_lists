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
	fmt.Println()

	RemoveDuplicates(&head)

	fmt.Print("\ndeduped list: ")
	list.Print(head)
	fmt.Println()

}

func RemoveDuplicates(source **list.Node) {

	for n := *source; n != nil; n = n.Next {
		if m := n.Next; m != nil {
			if n.Data == m.Data {
				n.Next = m.Next
			}
		}
	}
}
