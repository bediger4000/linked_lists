package list

import "strconv"

// MultiCompose gets []string, where the elements are string
// representations of integers, with occasional "--" elements.
// Each "--" divides a linked list from the next list.
func MultiCompose(stringValues []string) (lists []*Node) {

	if len(stringValues) == 0 {
		return
	}

	var head, tail *Node

	for _, str := range stringValues {

		if str == "--" {
			if head != nil {
				tail.Next = nil
				lists = append(lists, head)
				head = nil
			}
			continue
		}

		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}

		if head == nil {
			head = &Node{Data: num}
			tail = head
			continue
		}

		tail.Next = &Node{Data: num}
		tail = tail.Next
	}

	if head != nil {
		tail.Next = nil
		lists = append(lists, head)
	}

	return
}

// Compose a linked list from a slice of strings,
// which should be representations of integers
func Compose(stringValues []string) (head *Node) {
	if len(stringValues) == 0 {
		return
	}
	var tail *Node
	for _, str := range stringValues {
		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		if head == nil {
			head = &Node{Data: num}
			tail = head
			continue
		}
		tail.Next = &Node{Data: num}
		tail = tail.Next
	}
	return
}
