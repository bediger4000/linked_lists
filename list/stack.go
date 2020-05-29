package list

// Push puts a value on the top of the stack (FIFO)
func Push(stack **Node, value int) {
	tmp := &Node{Data: value}
	tmp.Next = *stack
	*stack = tmp
}

// Pop removes a value from the top of the stack
// and gives it back
func Pop(stack **Node) int {
	if *stack == nil {
		return 0
	}

	tmp := *stack
	*stack = (*stack).Next
	return tmp.Data
}

// Empty returns true when stack has no values,
// false otherwise
func Empty(stack **Node) bool {
	if *stack == nil {
		return true
	}
	return false
}
