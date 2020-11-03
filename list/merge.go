package list

// SortedMerge assumes formal parameters are numericall
// sorted, ascending order linked lists, returns a merged
// linked list. Formal parameter lists are subsumed
// in merged list, no longer are separate entities.
func SortedMerge(list1, list2 *Node) (head *Node) {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	listp := &list1
	if list2.Data > list1.Data {
		listp = &list2
	}
	head = *listp
	*listp = (*listp).Next
	head.Next = nil
	tail := head

	for {
		if list1 == nil {
			tail.Next = list2
			break
		}
		if list2 == nil {
			tail.Next = list1
			break
		}

		listp := &list1
		if list2.Data > list1.Data {
			listp = &list2
		}
		tail.Next = *listp
		tail = tail.Next
		*listp = (*listp).Next
		tail.Next = nil
	}

	return
}
