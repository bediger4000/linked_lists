package main

import (
	"fmt"
	"os"

	"linked_lists/dllist"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Determine if a doubly-linked list is palindromic\n")
		fmt.Fprintf(os.Stderr, "Usage: %s N0 [N1 N2 N3 ...]\n", os.Args[0])
		os.Exit(1)
	}

	head := dllist.Compose(os.Args[1:])

	fmt.Printf("Original list:\n")
	dllist.Print(head)
	fmt.Println()

	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}

	fmt.Printf("Original list in reverse:\n")
	for node := tail; node != nil; node = node.Prev {
		fmt.Printf("%d -> ", node.Data)
	}
	fmt.Println()

	if isPalindrome(head) {
		fmt.Printf("It is palindromic\n")
		return
	}

	fmt.Printf("It is NOT palindromic\n")
}

// isPalindrome returns true if the input list is "palindromic",
// has the same element values going both ways along the list,
// until the middle node, if one exists
func isPalindrome(head *dllist.Node) bool {
	if head == nil {
		return true
	}

	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}

	for ; head != tail; head, tail = head.Next, tail.Prev {
		if head.Data != tail.Data {
			return false
		}
	}

	return true
}
