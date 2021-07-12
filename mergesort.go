package main

import (
	"linked_lists/list"
	"os"
)

func main() {
	head := list.Compose(os.Args[1:])
	list.Print(head)
}
