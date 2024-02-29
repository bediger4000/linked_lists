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
	fmt.Printf("Cycle meets at node value %d\n", linkValue)

	head := list.Compose(os.Args[2:])
	list.Print(head)
	fmt.Println("")
	count := 0
	for n := head; n != nil; n = n.Next {
		count++
	}
	fmt.Printf("%d items in list\n", count)

	if link := list.FindItem(head, linkValue); link == nil {
		fmt.Printf("Could not find list node value %d\n", linkValue)
		fmt.Printf("no cycle introduced\n")
	} else {
		tail := head
		for tail.Next != nil {
			tail = tail.Next
		}

		tail.Next = link
		list.PrintN(head, count+3)
		fmt.Println("")
	}

	if cycleExists, at, period := list.Floyds(head); cycleExists {
		fmt.Printf("Cycle of period %d exists, meets at node value %d\n",
			period, at.Data)
	} else {
		fmt.Println("no cycle exists")
	}

}
func usage() {
	fmt.Printf("Find if a linked list contains a cycle\n")
	fmt.Printf("Invocation:\n./floyds N a b c N d e f g h\n")
	fmt.Printf("N is the value of the node that's the head of the cycle\n")
	fmt.Printf("All command line values string reps of integers\n")
}
