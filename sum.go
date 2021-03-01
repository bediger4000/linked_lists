package main

/*
This problem was asked by Microsoft.

Let's represent an integer in a linked list format by having each
node represent a digit in the number.
The nodes make up the number in reversed order.

For example, the following linked list:

1 -> 2 -> 3 -> 4 -> 5

is the number 54321.

Given two linked lists in this format,
return their sum in the same linked list format.

For example, given

9 -> 9

5 -> 2

return 124 (99 + 25) as:

4 -> 2 -> 1
*/

import (
	"fmt"
	"os"

	"linked_lists/list"
)

func main() {

	var list1, list2 *list.Node

	for idx, str := range os.Args[0:] {
		if idx == 0 {
			continue
		}
		if str == "--" {
			list1 = list.Compose(os.Args[1:idx])
			list2 = list.Compose(os.Args[idx+1:])
		}
	}

	list.Print(list1)
	fmt.Println()
	list.Print(list2)
	fmt.Println()

	sum := add(list1, list2)
	list.Print(sum)
	fmt.Println()
}

func add(list1, list2 *list.Node) *list.Node {
	var tail, result *list.Node
	carry := 0

	for {
		if list1 == nil || list2 == nil {
			break
		}
		next := &list.Node{Data: list1.Data + list2.Data + carry}
		carry = 0
		if next.Data > 9 {
			carry = 1
			next.Data %= 10
		}
		list1 = list1.Next
		list2 = list2.Next
		if result == nil {
			result = next
		} else {
			tail.Next = next
		}
		tail = next
	}

	remaining := list1
	if remaining == nil {
		remaining = list2
	}

	for remaining != nil {
		next := &list.Node{Data: remaining.Data + carry}
		carry = 0
		if next.Data > 9 {
			carry = 1
			next.Data %= 10
		}
		remaining = remaining.Next
		if result == nil {
			result = next
		} else {
			tail.Next = next
		}
		tail = next
	}

	if carry > 0 {
		tail.Next = &list.Node{Data: carry}
	}

	return result
}
