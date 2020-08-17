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

	kthlast := head
	lead := head
	for i := 0; i < k+2; i++ {
		lead = lead.Next
	}

	for lead != nil {
		kthlast = kthlast.Next
		lead = lead.Next
	}

	// Excise the node that is the k'th last
	kthlast.Next = kthlast.Next.Next

	fmt.Printf("Modified list:\n")
	list.Print(head)
	fmt.Println()
}
