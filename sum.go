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

// add constructs a linked list with Data values the
// sum of corresponding Data values from list1 and list2.
// Has to reverse the list it constructs because
func add(list1, list2 *list.Node) *list.Node {
	var stack *list.Node
	carry := 0

	for list1 != nil || list2 != nil {
		value := 0
		if list1 != nil {
			value += list1.Data
			list1 = list1.Next
		}
		if list2 != nil {
			value += list2.Data
			list2 = list2.Next
		}
		// New node constructed as a push onto a stack
		stack = &list.Node{Data: value + carry, Next: stack}
		carry = 0
		if stack.Data > 9 {
			carry = 1
			stack.Data %= 10
		}
	}

	if carry > 0 {
		stack = &list.Node{Data: carry, Next: stack}
	}

	// reverse the stack into the order the problem
	// statement wants.
	var result *list.Node
	for stack != nil {
		tmp := stack.Next
		stack.Next = result
		result = stack
		stack = tmp
	}

	return result
}
