package main

/*
Given two singly linked lists that intersect at some point,
find the intersecting node.
The lists are non-cyclical.

For example,
given A = 3 -> 7 -> 8 -> 10 and B = 99 -> 1 -> 8 -> 10,
return the node with value 8.

In this example,
assume nodes with the same value are the exact same node objects.

Do this in O(M + N) time (where M and N are the lengths of the lists)
and constant space.
*/

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"linked_lists/list"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("intersection of 2 linked lists\n")
		fmt.Printf("usage: %s <intersectionvalue> n1.0 ... -- n2.0 ...\n", os.Args[0])
		return
	}

	intersectionValue, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Insecting lists at node value %d\n", intersectionValue)

	lists := list.MultiCompose(os.Args[2:])

	if len(lists) > 2 {
		log.Fatal("require only 2 lists\n")
	}

	fmt.Printf("Found %d lists\n", len(lists))

	for _, l := range lists {
		list.Print(l)
		fmt.Println()
	}

	if !createIntersection(intersectionValue, lists) {
		fmt.Printf("lists did not intersect\n")
		return
	}

	fmt.Printf("After creating an intersection:\n")
	for _, l := range lists {
		list.Print(l)
		fmt.Println()
	}

	intersectingNode := findIntersection(lists[0], lists[1])

	if intersectingNode == nil {
		fmt.Printf("Did not find intersection\n")
		return
	}

	fmt.Printf("Intersecting node at %p, value %d\n", intersectingNode, intersectingNode.Data)
}

// findIntersection
func findIntersection(hd1, hd2 *list.Node) *list.Node {
	len1 := listLength(hd1)
	len2 := listLength(hd2)

	for len1 > len2 {
		hd1 = hd1.Next
		len1--
	}

	for len2 > len1 {
		hd2 = hd2.Next
		len2--
	}

	var intersection *list.Node

	for n1, n2 := hd1, hd2; n1 != nil && n2 != nil; n1, n2 = n1.Next, n2.Next {
		if n1 == n2 {
			intersection = n1
			break
		}
	}

	return intersection
}

func listLength(head *list.Node) int {
	m := 0
	for p := head; p != nil; p = p.Next {
		m++
	}
	return m
}

// createIntersection puts the 2 lists (lists[0] and lists[1])
// together at a node in each with value intersectionValue
func createIntersection(intersectionValue int, lists []*list.Node) bool {
	hd0 := lists[0]
	hd1 := lists[1]

	indir0 := &hd0

	for (*indir0).Data != intersectionValue {
		indir0 = &(*indir0).Next
		if *indir0 == nil {
			break
		}
	}

	if *indir0 == nil {
		return false
	}

	indir1 := &hd1

	for (*indir1).Data != intersectionValue {
		indir1 = &(*indir1).Next
		if *indir1 == nil {
			break
		}
	}

	if *indir1 == nil {
		return false
	}

	*indir0 = *indir1

	return true
}
