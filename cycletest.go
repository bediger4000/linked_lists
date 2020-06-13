package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"linked_lists/list"
)

func main() {
	if len(os.Args) == 1 {
		usage()
		return
	}
	linkValue, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	head := list.Compose(os.Args[2:])

	link := list.FindItem(head, linkValue)

	if link == nil {
		log.Fatalf("Could not find list node value %d\n", linkValue)
	}

	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = link

	if list.ContainsCycle1(head) {
		fmt.Printf("Cycle1 algorithm returns true\n")
	}
	if list.ContainsCycle2(head) {
		fmt.Printf("Cycle2 algorithm returns true\n")
	} else {
		log.Fatal(fmt.Errorf("list does not contain a cycle"))
	}

	mtg := list.CycleHead1(head)
	fmt.Printf("1. Cycle at node value %d\n", mtg.Data)
	mtg = list.CycleHead2(head)
	fmt.Printf("2. Cycle at node value %d\n", mtg.Data)
}
func usage() {
	fmt.Printf("Find if a linked list contains a cycle\n")
	fmt.Printf("Invocation:\n./cycletest N a b c N d e f g h N\n")
	fmt.Printf("N is the value of the node that's the head of the cycle\n")
	fmt.Printf("All command line values string reps of integers\n")
}
