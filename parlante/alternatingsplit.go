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

	var aRef, bRef *list.Node
	AlternatingSplit(&head, &aRef, &bRef)

	fmt.Println("a list: ")
	list.Print(aRef)
	fmt.Println("\nb list: ")
	list.Print(bRef)
	fmt.Println()
}

func AlternatingSplit(source, aRef, bRef **list.Node) {

	for n := *source; n != nil; {
		tmp := n.Next
		n.Next = *aRef
		*aRef = n
		n = tmp
		if n != nil {
			tmp = n.Next
			n.Next = *bRef
			*bRef = n
			n = tmp
		}
	}
}
