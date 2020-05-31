package main

import (
	"fmt"
	"os"

	"linked_lists/list"
)

func main() {

	var list1 *list.Node

	list1 = list.Compose(os.Args[1:])

	list.Print(list1)
	fmt.Println()

	reversed := list.Reverse(list1)
	list.Print(reversed)
	fmt.Println()
}
