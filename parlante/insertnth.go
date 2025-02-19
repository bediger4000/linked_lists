package main

import (
	"fmt"
	"linked_lists/list"
	"log"
	"os"
	"strconv"
)

func main() {
	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if index < 0 {
		log.Fatal("index must be >= 0")
	}

	value, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	head := list.Compose(os.Args[3:])
	list.Print(head)
	fmt.Println()

	InsertNth(&head, index, value)

	list.Print(head)
	fmt.Println()

}

func InsertNth(head **list.Node, index int, value int) {

	node := head

	for i := 0; i < index && *node != nil; i++ {
		node = &((*node).Next)
	}

	newNode := &list.Node{
		Data: value,
		Next: *node,
	}

	*node = newNode
}
