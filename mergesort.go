package main

import (
	"fmt"
	"linked_lists/list"
	"os"
)

func main() {
	listLength := len(os.Args[1:])
	head := list.Compose(os.Args[1:])
	list.PrintN(head, listLength+1)
	fmt.Println()

	nl := mergesort(head)
	fmt.Println("Merged:")
	list.PrintN(nl, listLength+1)
	fmt.Println()
}

func mergesort(head *list.Node) *list.Node {

	if head == nil {
		return nil
	}

	// clever way to append to the tail of a list:
	// first time, set hd and tl, and reset appendTl
	// pointer to realappend, which just appends to a
	// (already initialized) linked list tail pointer.
	var appendTl func(n *list.Node)
	var hd, tl *list.Node
	realappend := func(n *list.Node) {
		tl.Next = n
		tl = n
	}
	startAppend := func(n *list.Node) {
		hd = n
		tl = n
		appendTl = realappend
	}
	appendTl = startAppend

	p := head
	mergecount := 2 // just to pass the first for-test

	for k := 1; mergecount > 1; k *= 2 {

		mergecount = 0

		for p != nil {

			psize := 0
			q := p
			for i := 0; q != nil && i < k; i++ {
				psize++
				q = q.Next
			}

			qsize := psize

			for psize > 0 && qsize > 0 && q != nil {
				if p.Data < q.Data {
					appendTl(p)
					p = p.Next
					psize--
					continue
				}
				appendTl(q)
				q = q.Next
				qsize--
			}

			for ; psize > 0 && p != nil; psize-- {
				appendTl(p)
				p = p.Next
			}

			for ; qsize > 0 && q != nil; qsize-- {
				appendTl(q)
				q = q.Next
			}

			p = q

			mergecount++
		}

		p = hd
		head = hd

		hd = nil
		tl.Next = nil
		tl = nil
		appendTl = startAppend
	}

	return head
}
