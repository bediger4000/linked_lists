package main

import (
	"fmt"
	"linked_lists/list"
	"os"
)

func main() {

	head := list.Compose(os.Args[1:])
	fmt.Print("source list: ")
	list.Print(head)
	fmt.Println()

	var frontRef, backRef *list.Node

	FrontBackSplit(&head, &frontRef, &backRef)

	fmt.Print("\nsource list: ")
	list.Print(head)
	fmt.Print("\nFront list: ")
	list.Print(frontRef)
	fmt.Print("\nBack list:  ")
	list.Print(backRef)
	fmt.Println()

}

// FrontBackSplit1 - for lists of length 2n+1, the "front" list
// has n elements, the "back" list has n+1
func FrontBackSplit(source, frontRef, backRef **list.Node) {
	rabbit, turtle := *source, source
	for rabbit != nil {
		turtle = &(*turtle).Next
		if rabbit = rabbit.Next; rabbit != nil {
			rabbit = rabbit.Next
		}
	}
	*backRef = *turtle
	*turtle = nil
	*frontRef = *source
	*source = nil
}

// FrontBackSplit - for lists of length 2n+1, the "front" list
// has n+1 elements, the "back" list has n
func FrontBackSplit1(source, frontRef, backRef **list.Node) {
	rabbit, turtle := (*source).Next, source
	for rabbit != nil {
		turtle = &(*turtle).Next
		if rabbit = rabbit.Next; rabbit != nil {
			rabbit = rabbit.Next
		}
	}
	*backRef = *turtle
	*turtle = nil // terminate front list
	*frontRef = *source
	*source = nil
}
