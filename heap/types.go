// Package heap implements a min-heap priority queue specifically
// for the problem of creating a single, sorted, singly-linked
// list from k sorted, singly-linked lists.
//
// Heap nodes are heads of singly-linked lists,
// from package list in this repo.
//
// Pops the head node from the minimum-valued singly-linked
// list when the Delete operation occurs. Maintains partial order
// based on value of head-of-list nodes' data.
// Heaps are implemented as variable-length arrays.
package heap

import "linked_lists/list"

// Heap as an array: simplifies maintaining the shape
// of the heap.
type Heap []*list.Node

/*
Thus the children of the node at position n would
2n + 1 and 2n + 2 in a zero-based array.
Computing the index of the parent node of n-th element is also
straightforward.
Similarly, for zero-based arrays, is the parent is
located at position (n-1)/2 (floored).
*/
