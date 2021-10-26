package main

import (
	"fmt"
	"linked_lists/list"
	"os"
)

/*

This problem was asked by Amazon.

Given a linked list, remove all consecutive nodes that sum to zero.
Print out the remaining nodes.

For example,
suppose you are given the input 3 -> 4 -> -7 -> 5 -> -6 -> 6.
In this case, you should first remove 3 -> 4 -> -7, then -6 -> 6, leaving only 5.

*/

func main() {

	head := list.Compose(os.Args[1:])

	var newhead, newtail *list.Node
	appnd := func(n *list.Node) {
		if newhead == nil {
			newhead = n
			newtail = n
			newtail.Next = nil
			return
		}
		newtail.Next = n
		newtail = n
		newtail.Next = nil
	}

	for hd := head; hd != nil; {
		currentSum := 0
		var node *list.Node

		for node = hd; node != nil; node = node.Next {
			currentSum += node.Data
			if currentSum == 0 {
				fmt.Printf("Delete nodes from %d to %d\n", hd.Data, node.Data)
				break
			}
		}

		if node != nil {
			hd = node.Next
		} else {
			tmp := hd.Next
			appnd(hd)
			hd = tmp
		}
	}

	list.Print(newhead)
	fmt.Println()
}
