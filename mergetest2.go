package main

// merge test benchmark using crypto/rand for the sort keys.
// Checking if math/rand accidentally created "adversarial"
// sort keys.

import (
	"crypto/rand"
	"fmt"
	"linked_lists/list"
	"log"
	"math"
	"math/big"
	"os"
	"time"
)

func main() {

	// rand.Seed(time.Now().UnixNano() | int64(os.Getpid()))

	for n := 10000; n < 8000000; n += 200000 {
		var total time.Duration
		for i := 0; i < 10; i++ {
			before := time.Now()
			head := randomValueList(n)

			before = time.Now()
			nl := mergesort(head)
			elapsed := time.Since(before)

			if !isSorted(nl) {
				os.Exit(1)
			}
			total += elapsed

		}
		total /= 10.0
		fmt.Printf("%d\t%.04f\n", n, total.Seconds())
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

var max = big.NewInt(math.MaxInt32)

func randomValueList(n int) *list.Node {

	var head *list.Node

	for i := 0; i < n; i++ {
		mp, err := rand.Int(rand.Reader, max)
		if err != nil {
			log.Fatal(err)
		}
		ui := mp.Int64()

		head = &list.Node{
			Data: int(ui),
			Next: head,
		}
	}

	return head
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
