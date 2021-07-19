package main

import (
	"fmt"
	"linked_lists/heap"
	"linked_lists/list"
	"os"
)

func main() {

	heads := list.MultiCompose(os.Args[1:])

	var nl *list.Node
	// each list has to be sorted
	nl = heapMerge(heads)
	fmt.Println("Merged:")
	list.Print(nl)
	fmt.Println()

	if !isSorted(nl) {
		fmt.Printf("NOT SORTED IN ASCENDING ORDER\n")
	}
}

func isSorted(head *list.Node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	for ; head.Next != nil; head = head.Next {
		if head.Data > head.Next.Data {
			return false
		}
	}
	return true
}

func heapMerge(heads []*list.Node) *list.Node {
	var h heap.Heap
	for _, hd := range heads {
		h = h.Insert(hd)
	}
	var hd, tl, tmp *list.Node

	for len(h) > 0 {
		h, tmp = h.Delete()
		if hd == nil {
			hd = tmp
			tl = tmp
			continue
		}
		tl.Next = tmp
		tl = tmp
	}

	return hd
}
