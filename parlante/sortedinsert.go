package main

import (
	"fmt"
	"linked_lists/list"
	"log"
	"os"
	"strconv"
)

func main() {
	newvalue, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	head := list.Compose(os.Args[2:])
	list.Print(head)
	fmt.Println()

	SortedInsert(&head, &list.Node{Data: newvalue})

	list.Print(head)
	fmt.Println()

}

func SortedInsert(head **list.Node, newnode *list.Node) {

	node := head

	for *node != nil && (*node).Data < newnode.Data {
		node = &((*node).Next)
	}

	newnode.Next = *node
	(*node) = newnode
}
