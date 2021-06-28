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

	var tmp, hd, tl *list.Node

	for current := head; current != nil; current = tmp {

		tmp = current.Next
		current.Next = nil

		if hd == nil {
			hd = current
			tl = current
			continue
		}

		if current.Data < k {
			current.Next = hd
			hd = current
		}

		if current.Data >= k {
			tl.Next = current
			tl = current
		}
	}

	return hd
}
