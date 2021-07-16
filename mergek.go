package main

import (
	"fmt"
	"linked_lists/list"
	"os"
)

func main() {

	heads := list.MultiCompose(os.Args[1:])

	// each list has to be sorted
	nl := mergek(heads)
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

func mergek(heads []*list.Node) *list.Node {

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

	combined := heads[0]

	for _, head := range heads[1:] {

		for combined != nil && head != nil {
			if combined.Data < head.Data {
				append(combined)
				combined = combined.Next
				continue
			}
			append(head)
			head = head.Next
		}

		if combined != nil {
			tl.Next = combined
		}
		if head != nil {
			tl.Next = head
		}

		combined = hd
		hd = nil
		tl = nil
	}

	return combined
}
