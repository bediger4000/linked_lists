// Package list implements a doubly-linked list of integer-valued nodes,
// along with input and output functions.
package dllist

// Node is an element of a doubly-linked list
type Node struct {
	Data int
	Next *Node
	Prev *Node
}
