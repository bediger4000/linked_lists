package main

import (
	"fmt"
	"linked_lists/list"
	"os"
	"strconv"
)

func main() {
	var idx int
	var str string
	q := &list.Queue{}
	for idx, str = range os.Args[1:] {
		if str == "--" {
			break
		}
		val, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		q.Enqueue(val)
	}

	for !q.Empty() {
		val := q.Dequeue()
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	if str != "--" {
		return
	}

	for _, str = range os.Args[idx+1:] {
		val, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		q.Enqueue(val)
	}

	for !q.Empty() {
		val := q.Dequeue()
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
