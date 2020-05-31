// Package list implements a linked list of integer-valued nodes,
// along with input and output functions.
package list

// Node is an element of a linked list
type Node struct {
	Data int
	Next *Node
}

// Stack type to segregate LIFO functions
type Stack Node
