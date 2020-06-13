package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"linked_lists/list"
)

func main() {

	desired, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Remove list element containing '%c'\n", desired)

	head := list.Compose(os.Args[2:])

	list.Print(head)
	fmt.Println()

	node := list.FindItem(head, desired)

	switch node {
	case nil:
		fmt.Printf("Did not find '%d' in list\n", desired)
		return
	default:
		fmt.Printf("Element at '%p' contains '%d'\n", node, node.Data)
	}

	list.Remove(node)

	fmt.Printf("list without '%d':\n", desired)
	list.Print(head)
	fmt.Println()
}

// Remove a node from a list by copying the next node to the
// current node. A little wasteful.
func Remove(node *Node) {
	for node != nil {
		tmp := node.Next
		if tmp != nil {
			node.Data = tmp.Data
			node.Next = tmp.Next
		}
		node = node.Next
	}
}
