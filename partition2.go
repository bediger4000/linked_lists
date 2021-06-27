package main

import (
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

	head = partition(k, head)

	list.Print(head)
}

func partition(k int, head *list.Node) *list.Node {
	var tmp, beforeHead, afterHead *list.Node

	for current := head; current != nil; current = tmp {
		tmp = current.Next
		current.Next = nil

		if current.Data < k {
			current.Next = beforeHead
			beforeHead = current
		}

		if current.Data >= k {
			current.Next = afterHead
			afterHead = current
		}
	}

	head = beforeHead
	if head == nil {
		// no list items data < the parition value
		head = afterHead
	} else if afterHead != nil {
		// list items with data < the partition value,
		// and list items with data >= the parition value.
		// Find tail of list, then append the ">=" sub-list
		tmp = head
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		if tmp != nil {
			// not a zero-length input list
			tmp.Next = afterHead
		}
	}

	return head
}
