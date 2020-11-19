package main

import (
	"fmt"
	"os"

	"linked_lists/dllist"
)

func main() {
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
