package main

import (
	"fmt"
	"os"
	"strconv"
	"unsafe"
)

// XorNode instances make up an XOR-list, the conceptual doubly-linked
// list of the problem statement. Go's type safety made this harder than
// it would be in C. The "both" element should constitute an Xor of the
// addresses of the node previous and the node next in list from the
// XorNode instance in question.
type XorNode struct {
	data int
	both uintptr
}

// add puts a new node on the other end of the list from head.
// Use it like: head,tail = add(head, newnode)
// Returns head and tail of the XOR-list in defiance of the spec
// so that I can see the XOR-list work forwards and backwards
func add(head *XorNode, element *XorNode) (*XorNode, *XorNode) {
	if head == nil {
		// element is the first element in list
		return element, element
	}
	if head.both == 0 {
		// head is the only element in the list
		head.both = uintptr(unsafe.Pointer(element))
		element.both = uintptr(unsafe.Pointer(head))
		return head, element
	}
	node := head
	var previous uintptr
	for {
		// Node "holds a field named both, which is an XOR of the next node
		// and the previous node".
		next := uintptr(unsafe.Pointer(node.both)) ^ previous
		previous = uintptr(unsafe.Pointer(node))
		if next == 0 {
			// node is the last element in list
			node.both = node.both ^ uintptr(unsafe.Pointer(element))
			// element.both xor of node and nil
			element.both = previous
			break
		}
		node = (*XorNode)(unsafe.Pointer(next))
	}

	return head, element
}

// get returns a node at index wantIndex (0-indexed!)
// based on the head node. Returns nil if the list is
// too short to have an element at wantIndex
func get(head *XorNode, wantIndex int) *XorNode {
	var previous uintptr
	for node, index := head, 0; true; index++ {
		if index == wantIndex {
			return node
		}
		next := uintptr(unsafe.Pointer(node.both)) ^ previous
		if next == 0 {
			break // end of list
		}
		previous = uintptr(unsafe.Pointer(node))
		node = (*XorNode)(unsafe.Pointer(next))
	}
	return nil // didn't get to wantIndex
}

// print writes a representation of the data in an XOR-list
// on stdout: 0 -> 1 -> 2 -> nil, or something like that,
// depends on the list.
func print(head *XorNode) {
	node := head
	var previous uintptr
	for {
		fmt.Printf("%d -> ", node.data)
		next := uintptr(unsafe.Pointer(node.both)) ^ previous
		previous = uintptr(unsafe.Pointer(node))
		if next == 0 {
			fmt.Println("nil")
			break
		}
		node = (*XorNode)(unsafe.Pointer(next))
	}
}

func main() {

	var head, tail *XorNode

	// Build the list from string representations of numbers
	// appearing on command line.
	nodeCount := 0
	for _, str := range os.Args[1:] {
		n, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		head, tail = add(head, &XorNode{data: n})
		nodeCount++
	}
	fmt.Printf("%d nodes in list\n", nodeCount)

	// Print the list forward and backward just to prove the Xor-property
	fmt.Printf("Front to back:\n")
	print(head)
	fmt.Printf("Back to front:\n")
	print(tail)

	// Retrieve every node in the list by its index
	for i := 0; i < nodeCount; i++ {
		node := get(head, i)
		fmt.Printf("node %d at %p, value %d\n", i, node, node.data)
	}

	node := get(head, nodeCount+10)
	if node == nil {
		fmt.Printf("Correctly did not find node at index %d\n", nodeCount+10)
	} else {
		fmt.Printf("Incorrectly found node at index %d, value %d\n", nodeCount+10, node.data)
	}
}
