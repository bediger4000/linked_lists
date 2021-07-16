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

	var hd, tl *list.Node
	append := func(n *list.Node) {
		if hd == nil {
			hd = n
			tl = n
			return
		}
		tl.Next = n
		tl = n
	}

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
					append(p)
					p = p.Next
					psize--
					continue
				}
				append(q)
				q = q.Next
				qsize--
			}

			for ; psize > 0 && p != nil; psize-- {
				append(p)
				p = p.Next
			}

			for ; qsize > 0 && q != nil; qsize-- {
				append(q)
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
	}

	return head
}
