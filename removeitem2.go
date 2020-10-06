package main

/*
https://medium.com/@bartobri/applying-the-linus-tarvolds-good-taste-coding-requirement-99749f37684a

Linus Torvald's example of "good taste" programming.
*/

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"linked_lists/list"
)

func main() {

	desired, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Remove list element containing '%c'\n", desired)

	head := list.Compose(os.Args[2:])

	list.Print(head)
	fmt.Println()

	indirect := &head

	for (*indirect).Data != desired {
		indirect = &(*indirect).Next
		if *indirect == nil {
			break
		}
	}

	if *indirect != nil {
		fmt.Printf("found list item with desired value %d\n", (*indirect).Data)
		*indirect = (*indirect).Next
	}

	fmt.Printf("list without '%d':\n", desired)
	list.Print(head)
	fmt.Println()
}
