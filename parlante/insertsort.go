package main

import (
	"fmt"
	"linked_lists/list"
	"os"
)

func main() {

	head := list.Compose(os.Args[1:])
	list.Print(head)
	fmt.Println()

	InsertSort(&head)

	list.Print(head)
	fmt.Println()

}

func InsertSort(head **list.Node) {

	var sortedHead *list.Node

	node := *head

	for node != nil {
		tmp := node
		node = node.Next
		SortedInsert(&sortedHead, tmp)
	}

	*head = sortedHead
}

func SortedInsert(head **list.Node, newnode *list.Node) {

	node := head

	for *node != nil && (*node).Data < newnode.Data {
		node = &((*node).Next)
	}

	newnode.Next = *node
	(*node) = newnode
}
