package list

import "fmt"

// Print runs a linked list and prints its values on stdout
func Print(list *Node) {
	for node := list; node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Data)
	}
}

// PrintN runs a linked list, printing first n values on stdout
func PrintN(list *Node, n int) {
	count := 0
	var node *Node
	for node = list; count < n && node != nil; node = node.Next {
		fmt.Printf("%d -> ", node.Data)
		count++
	}
	if node != nil {
		fmt.Print("...")
	}
}
