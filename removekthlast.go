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
	tmp := head
	for i := 0; i < k+2; i++ {
		tmp = tmp.Next
	}

	for tmp != nil {
		kthlast = kthlast.Next
		tmp = tmp.Next
	}

	kthlast.Next = kthlast.Next.Next
	fmt.Printf("Modified list:\n")
	list.Print(head)
	fmt.Println()
}

// isListPalindrome says true if the list constitutes a palindrome,
// false otherwise.
// 1. Find middle of list
// 2. Reverse list from middle to end
// 3. Check if two half-lists (one reversed) have same values in order.
func isListPalindrome(head *list.Node) bool {

	mid := middleNode(head)
	var tail *list.Node
	for tail = head; tail.Next != mid; tail = tail.Next {
	}
	tail.Next = nil
	fmt.Printf("First half:\n")
	list.Print(head)
	fmt.Printf("\n2nd half:\n")
	list.Print(mid)
	fmt.Println()

	reversed := list.Reverse(mid)
	fmt.Printf("2nd half reversed:\n")
	list.Print(reversed)
	fmt.Println()

	for head.Data == reversed.Data {
		head = head.Next
		reversed = reversed.Next
		if head == nil {
			break
		}
		if reversed == nil {
			break
		}
	}

	if head == nil && reversed == nil {
		return true
	}
	if head == nil {
		if reversed.Next == nil {
			return true
		}
	}
	return false
}

func middleNode(head *list.Node) *list.Node {
	slow := head
	fast := head

	// fast steps through list 2 at a time, slow 1 at a time.
	// when fast has nil value, slow points to center item.
	// even number of items in list cause a little trouble.
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		}
	}

	return slow
}
