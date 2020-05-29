package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"linked_lists/list"
)

func main() {
	var stack *list.Node

	fmt.Printf("Empty: %v\n", list.Empty(&stack))

	for _, str := range os.Args[1:] {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Println(err)
			continue
		}
		list.Push(&stack, num)
	}

	for !list.Empty(&stack) {
		value := list.Pop(&stack)
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}
