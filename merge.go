package main

import (
	"fmt"
	"os"

	"linked_lists/list"
)

func main() {

	var list1, list2 *list.Node

	for idx, str := range os.Args[0:] {
		if idx == 0 {
			continue
		}
		if str == "--" {
			list1 = list.Compose(os.Args[1:idx])
			list2 = list.Compose(os.Args[idx+1:])
		}
	}

	list.Print(list1)
	fmt.Println()
	list.Print(list2)
	fmt.Println()

	merged := list.SortedMerge(list1, list2)
	list.Print(merged)
	fmt.Println()
}
