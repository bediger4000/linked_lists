package main

import (
	"fmt"
	"os"

	"linked_lists/list"
)

func main() {
	lists := list.MultiCompose(os.Args[1:])

	fmt.Printf("Found %d lists\n", len(lists))

	for _, l := range lists {
		list.Print(l)
		fmt.Println()
	}
}
