package main

import (
	"fmt"
	"linked_lists/list"
	"os"
)

/*
 * Find the middle item in a linked list
 */

func main() {
	head := list.Compose(os.Args[1:])
	list.Print(head)
	fmt.Println()

	slow := head
	fast := head

	// fast steps through list 2 at a time, slow 1 at a time.
	// when fast has nil value, slow points to center item.
	// even number of items in list cause a little trouble.
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		}
	}

	fmt.Printf("Middle item has value %d\n", slow.Data)
}
