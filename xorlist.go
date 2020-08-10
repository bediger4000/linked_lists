package main

import (
	"fmt"
	"os"
	"strconv"
	"unsafe"
)

type XorNode struct {
	data int
	both uintptr
}

func add(root *XorNode, element *XorNode) *XorNode {
	fmt.Printf("root at %p, new node %p\n", root, element)
	if root == nil {
		return element
	}
	ptr := uintptr(unsafe.Pointer(root))
	node := root

	for node.both != 0 {
		ptr = ptr ^ node.both
		node = (*XorNode)(unsafe.Pointer(ptr))
	}

	tmp := uintptr(unsafe.Pointer(node)) ^ uintptr(unsafe.Pointer(element))
	node.both = uintptr(unsafe.Pointer(tmp))
	element.both = 0

	return root
}

func get(root *XorNode, index int) *XorNode {
	ptr := uintptr(unsafe.Pointer(root))
	node := root

	for i := 0; i < index; i++ {
		ptr = ptr ^ uintptr(unsafe.Pointer(node.both))
		prevnode := node
		node = (*XorNode)(unsafe.Pointer(ptr))
		if prevnode == node {
			break
		}
	}

	return node
}

func print(node *XorNode) {

	ptr := uintptr(unsafe.Pointer(node))

	for node != nil {
		fmt.Printf("ptr %x\n", ptr)
		fmt.Printf("node at %p\n", node)
		fmt.Printf("node.data %d\n", node.data)
		fmt.Printf("node.both %x\n", node.both)

		prevnode := node
		ptr = ptr ^ uintptr(unsafe.Pointer(node.both))
		node = (*XorNode)(unsafe.Pointer(ptr))
		fmt.Printf("next node at %p\n", node)
		if node == prevnode {
			break
		}
	}
	fmt.Println()
}

func main() {

	var head, tail *XorNode

	nodeCount := 0
	for _, str := range os.Args[1:] {
		n, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		tail = &XorNode{data: n}
		head = add(head, tail)
		nodeCount++
	}
	fmt.Printf("%d nodes in list\n", nodeCount)
	print(head)
	fmt.Printf("Back to front:\n")
	print(tail)

	for i := 0; i < nodeCount; i++ {
		node := get(head, i)
		fmt.Printf("node %d at %p, value %d\n", i, node, node.data)
	}
}
