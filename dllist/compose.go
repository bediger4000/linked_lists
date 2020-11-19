package dllist

import "strconv"

// Compose a doubly-linked list from a slice of strings,
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
		node := &Node{Data: num}
		tail.Next = node
		node.Prev = tail
		tail = node
	}
	return
}
