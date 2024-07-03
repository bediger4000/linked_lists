package main

import (
	crand "crypto/rand"
	"flag"
	"fmt"
	"linked_lists/list"
	"log"
	"math"
	"math/big"
	"math/rand"
	"os"
	"time"
)

func main() {
	useCryptoRand := flag.Bool("c", false, "use cryptographic PRNG")
	composeSequential := flag.Bool("s", false, "compose a sequential list")
	n := flag.Int("n", 99, "number of integer-value nodes in list")
	flag.Parse()

	rand.Seed(time.Now().UnixNano() | int64(os.Getpid()))

	var head *list.Node
	if *composeSequential {
		head = sequentialValueList(*n)
	} else {
		head = randomValueList(*n, *useCryptoRand)
	}
	list.Print(head)
	fmt.Println()

	newHead := recursiveMergeSort(head)

	if !isSorted(newHead) {
		fmt.Printf("list is not sorted correctly\n")
	}

	fmt.Printf("%d nodes in sorted list\n", listSize(newHead))
	list.Print(newHead)
	fmt.Println()
}

func listSize(head *list.Node) int {
	sz := 0
	for ; head != nil; head = head.Next {
		sz++
	}
	return sz
}

var maxInt = big.NewInt(math.MaxInt32)

func randomValueList(n int, useCheapRand bool) *list.Node {

	var head *list.Node

	for i := 0; i < n; i++ {
		var ri int
		if useCheapRand {
			ri = rand.Int()
		} else {
			mp, err := crand.Int(crand.Reader, maxInt)
			if err != nil {
				log.Fatal(err)
			}
			ri = int(mp.Int64())
		}
		head = &list.Node{
			Data: ri,
			Next: head,
		}
	}

	return head
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

func recursiveMergeSort(head *list.Node) *list.Node {
	if head.Next == nil {
		// single node list is sorted by definiton
		return head
	}

	// because of recursion bottoming out at a 1-long-list,
	// head points to a list of at least 2 elements.

	// Setting rabbit and turtle like this means we split an
	// odd-length-list (head) into lists of length n (right)
	// and n+1 (left).
	rabbit, turtle := head.Next, &head

	for rabbit != nil {
		turtle = &(*turtle).Next
		if rabbit = rabbit.Next; rabbit != nil {
			rabbit = rabbit.Next
		}
	}

	right := *turtle
	*turtle = nil

	left := recursiveMergeSort(head)
	right = recursiveMergeSort(right)

	var h, t *list.Node
	if left.Data < right.Data {
		h, t = left, left
		left = left.Next
	} else {
		h, t = right, right
		right = right.Next
	}

	// left and right are either equal in length, or right is one
	// node longer, but the "<" check might take more from one list
	// than the other. Have to check both for nil.
	for left != nil && right != nil {
		var n *list.Node
		if left.Data < right.Data {
			n = left
			left = left.Next
		} else {
			n = right
			right = right.Next
		}
		t.Next = n
		t = t.Next
		// At the end of this for-loop, t.Next ends up being nil
		// because of the left/right list splitting.
	}

	// Either left or right are nil. If left == nil,
	// assigning nil to t.Next is no issue.
	t.Next = left
	if right != nil {
		// but if right is nil, can't assign nil to t.Next,
		// because left was non-nil.
		t.Next = right
	}

	return h
}

func sequentialValueList(sz int) *list.Node {
	var head *list.Node

	for i := 1; i <= sz; i++ {
		n := &list.Node{Data: i, Next: head}
		head = n
	}

	return head
}
