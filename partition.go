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
	var tmp, beforeHead, afterHead, beforeTail, afterTail *list.Node

	for current := head; current != nil; current = tmp {
		tmp = current.Next
		current.Next = nil

		if current.Data < k {
			if beforeHead == nil {
				beforeHead = current
				beforeTail = current
				continue
			}
			beforeTail.Next = current
			beforeTail = current
			continue
		}

		if current.Data >= k {
			if afterHead == nil {
				afterHead = current
				afterTail = current
				continue
			}
			afterTail.Next = current
			afterTail = current
			continue
		}
	}

	if beforeTail != nil {
		beforeTail.Next = afterHead
	}

	head = beforeHead
	if head == nil {
		head = afterHead
	}

	return head
}
