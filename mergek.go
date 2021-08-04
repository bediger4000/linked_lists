package main

// Two methods of combining K sorted singly linked lists into
// a single sorted singly linked list.
//
// ./merge 1 2 6 -- 0 3 7 -- 4 5 8 10
// Merged:
// 0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 10 ->

import (
	"flag"
	"fmt"
	"linked_lists/list"
	"os"
)

func main() {

	linearCombine := flag.Bool("k", false, "combine one at a time")
	flag.Parse()

	heads := list.MultiCompose(os.Args[1:])

	var nl *list.Node
	// each list has to be sorted
	if *linearCombine {
		nl = mergek(heads)
	} else {
		nl = recursiveMerge(heads)
	}
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

// mergek creates a single, numerically sorted singly-linked
// list from its argument slice of sorted singly-linked lists.
// Iteratively merge the lists into a singly combined list.
// Destroys lists in the argument slice.
func mergek(heads []*list.Node) *list.Node {

	combined := heads[0]

	for _, head := range heads[1:] {
		combined = merge2(combined, head)
	}

	return combined
}

// recursiveMerge emulates a merge sort, only on a slice
// of already sorted singly-linked lists. "Emulates" in the
// sense that it doesn't check lengths of those lists, so there's
// not power-of-2 size chunks of things to merge together.
func recursiveMerge(heads []*list.Node) *list.Node {
	if len(heads) == 1 {
		return heads[0]
	}
	l := len(heads) / 2
	left := recursiveMerge(heads[:l])
	right := recursiveMerge(heads[l:])
	return merge2(left, right)
}

// merge2 creates a sorted singly-linked list
// by destructively combining 2 input sorted linked lists.
func merge2(p, q *list.Node) *list.Node {

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

	for p != nil && q != nil {
		if p.Data < q.Data {
			append(p)
			p = p.Next
			continue
		}
		append(q)
		q = q.Next
	}

	if p != nil {
		tl.Next = p
	}
	if q != nil {
		tl.Next = q
	}

	return hd
}
