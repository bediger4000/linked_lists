package main

import (
	"fmt"
	"linked_lists/list"
	"log"
	"os"
	"strconv"
)

/*

Given a linked list and a positive integer k, rotate the list to the right by k places.

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

// rotate finds the kth node in the list,
// which will be returned as the head of a
// rotated list.
// Find the tail of the list (which isn't rotated yet),
// set tail.Next to head.
// Find the node before the kth node - it's the new tail
// node. Set its Next element to nil.
func rotate(k int, head *list.Node) *list.Node {

	// Find the kth node in the list, or run off the end
	// of the list, if it's shorter than k elements
	kth := head
	for i := 0; i < k && kth != nil; i++ {
		kth = kth.Next
	}
	if kth == nil {
		fmt.Fprintf(os.Stderr, "List less than %d elements long\n", k)
		return nil
	}

	// find tail node at end of original list
	var tail *list.Node
	for tail = kth; tail.Next != nil; tail = tail.Next {
	}

	// set tail.Next to old head node, find node before kth node
	for tail.Next = head; tail.Next != kth; tail = tail.Next {
	}
	// tail now the node before kth node
	tail.Next = nil

	return kth
}
