package main

import (
	"fmt"
	"linked_lists/list"
	"log"
	"os"
	"strconv"
)

func main() {
	index, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	head := list.Compose(os.Args[2:])

	n := GetNth(head, index)

	fmt.Printf("Found data value %d at index %d\n", n, index)
}

func GetNth(node *list.Node, n int) int {

	var value, length, i int

	for i = 0; node != nil && i <= n; node, i = node.Next, i+1 {
		length++
		if i == n {
			value = node.Data
			break
		}
	}

	if node == nil {
		// Got to end of list before getting to Nth node
		log.Fatalf("list length %d does not extend to desired index %d (%d)\n", length, n, i)
	}

	return value
}
