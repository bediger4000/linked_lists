package dllist

import "fmt"

// Print runs a linked list and prints its values on stdout
func Print(list *Node) {
	for node := list; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Data)
	}
}
