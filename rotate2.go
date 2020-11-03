package main

import (
	"fmt"
	"linked_lists/list"
	"log"
	"os"
	"strconv"
)

/*

Given a linked list and a positive integer k,
rotate the list to the right by k places.

For example,
given the linked list 7 -> 7 -> 3 -> 5 and k = 2,
it should become 3 -> 5 -> 7 -> 7.

Given the linked list 1 -> 2 -> 3 -> 4 -> 5 and k = 3,
it should become 3 -> 4 -> 5 -> 1 -> 2.

*/

func main() {
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	head := list.Compose(os.Args[2:])

	fmt.Print("Rotate linked list: ")
	list.Print(head)
	fmt.Printf(" by k = %d\n", k)
	newhead := rotate(k, head)
	if newhead != nil {
		fmt.Print("Rotated list:       ")
		list.Print(newhead)
		fmt.Println()
	}
}

// rotate creates a circular linked list, then
// walks it k places with both head and tail pointers,
// setting tail.Next to nil after that walk.
func rotate(k int, head *list.Node) *list.Node {

	// Find tail node at end of original list
	var tail *list.Node
	for tail = head; tail.Next != nil; tail = tail.Next {
	}

	// Make the list circular
	tail.Next = head

	// Move head and tail k steps through the list
	for i := 0; i < k; i++ {
		tail = tail.Next
		head = head.Next
	}

	// Break the circular list appropriately
	tail.Next = nil

	return head
}
