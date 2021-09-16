package main

/*
This problem was asked by Google.

Given a singly linked list and an integer k, remove the kth last element from
the list. k is guaranteed to be smaller than the length of the list.

The list is very long, so making more than one pass is prohibitively expensive.

Do this in constant space and in one pass.

*/

import (
	"fmt"
	"linked_lists/list"
	"log"
	"os"
	"strconv"
)

func main() {

	k, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	head := list.Compose(os.Args[2:])
	fmt.Printf("Original list:\n")
	list.Print(head)
	fmt.Println()

	head = removeKthLast(k, head)
	fmt.Printf("Modified list:\n")
	list.Print(head)
	fmt.Println()
}

func removeKthLast(k int, head *list.Node) *list.Node {
	leader := head

	// Advance past first k nodes in list.
	// Don't check for nil value of leader, k is less than length of list
	for i := 0; i < k; i++ {
		leader = leader.Next
	}

	// Linus Torvalds' "good taste" in programming method
	// indirect holds head-of-list address
	indirect := &head

	for leader = leader.Next; leader != nil; leader = leader.Next {
		indirect = &(*indirect).Next
	}

	// indirect holds address of Next field of k-1'th
	// node, or address of head. *indirect evaluates to the
	// address of the kth-last node

	(*indirect) = (*indirect).Next

	return head
}
