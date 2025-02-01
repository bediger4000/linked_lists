package main

import (
	"fmt"
	"linked_lists/list"
	"log"
	"os"
	"strconv"
)

func main() {
	given, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	head := list.Compose(os.Args[2:])

	n := Count(head, given)

	fmt.Printf("Found %d nodes with value %d\n", n, given)
}

func Count(head *list.Node, given int) int {
	count := 0

	for p := head; p != nil; p = p.Next {
		if p.Data == given {
			count++
		}
	}

	return count
}
