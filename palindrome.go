package main

/*
Note: Try to solve this task in O(n) time using O(1) additional space, where n
is the number of elements in l, since this is what you'll be asked to do during
an interview.

Given a singly linked list of integers, determine whether or not it's a
palindrome.

Example

For l = [0, 1, 0], the output should be
isListPalindrome(l) = true;

For l = [1, 2, 2, 3], the output should be
isListPalindrome(l) = false.

*/

import (
	"fmt"
	"linked_lists/list"
	"os"
)

func main() {

	head := list.Compose(os.Args[1:])
	fmt.Printf("Original list:\n")
	list.Print(head)
	fmt.Println()

	x := isListPalindrome(head)

	if x {
		fmt.Printf("It is palindromic\n")
	} else {
		fmt.Printf("It is NOT palindromic\n")
	}
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
