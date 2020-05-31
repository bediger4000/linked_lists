package list

// Reverse makes a list read backward. Messes up
// all pointers inputs list element is now last
// element in list.
func Reverse(head *Node) (reversed *Node) {
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = reversed
		reversed = tmp
	}
	return
}
