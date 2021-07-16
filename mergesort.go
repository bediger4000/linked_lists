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

	nl := merge2(head, listLength)
	fmt.Println("Merged:")
	list.PrintN(nl, listLength+1)
	fmt.Println()
}

func merge2(head *list.Node, listLength int) *list.Node {

	var hd, tl *list.Node
	append := func(n *list.Node) {
		fmt.Printf("Appending node with %d\n", n.Data)
		if hd == nil {
			hd = n
			tl = n
			return
		}
		tl.Next = n
		tl = n
	}

	p := head
	mergecount := 2
	k := 1

	for mergecount > 1 {

		mergecount = 0

		for p != nil {

			psize := 0
			q := p
			for i := 0; q != nil && i < k; i++ {
				psize++
				q = q.Next
			}
			fmt.Printf("psize %d, p at %d\n", psize, p.Data)
			if q != nil {
				fmt.Printf("q at %d\n", q.Data)
			} else {
				fmt.Println("q holds nil")
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

			fmt.Printf("1 psize %d, p at %d\n", psize, p.Data)
			if q != nil {
				fmt.Printf("1 qsize %d, q at %d\n", qsize, q.Data)
			}

			for ; psize > 0 && p != nil; psize-- {
				fmt.Printf("psize %d\n", psize)
				append(p)
				p = p.Next
			}

			for ; qsize > 0 && q != nil; qsize-- {
				fmt.Printf("qsize %d\n", qsize)
				append(q)
				q = q.Next
			}
			if p != nil {
				fmt.Printf("2 psize %d, p at %d\n", psize, p.Data)
			}
			if q != nil {
				fmt.Printf("2 qsize %d, q at %d\n", qsize, q.Data)
			}

			p = q

			mergecount++
		}

		p = hd
		head = hd
		tl.Next = nil
		hd = nil
		tl = nil

		k *= 2
	}

	return head
}
