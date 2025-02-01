package main

import (
	"fmt"
	"linked_lists/list"
	"log"
	"os"
)

func main() {

	head := list.Compose(os.Args[1:])

	list.Print(head)
	fmt.Println()

	data := Pop(&head)
	fmt.Printf("Head data was %d\n", data)

	list.Print(head)
	fmt.Println()

}

func Pop(head **list.Node) int {
	if *head == nil {
		log.Fatal("cannot pop empty list\n")
	}
	data := (*head).Data
	*head = (*head).Next
	return data
}
