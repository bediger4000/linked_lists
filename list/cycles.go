package list

// ContainsCycle1 constitutes a tortoise-and-hare cycle
// finder for a linked list.
func ContainsCycle1(head *Node) bool {

	tortoise := head
	hare := head

	for {
		if hare == nil {
			return false
		}
		hare = hare.Next
		if hare == nil {
			return false
		}
		hare = hare.Next
		tortoise = tortoise.Next
		if hare == tortoise {
			return true
		}
	}
	// Does it ever get here?
	return false
}

// ContainsCycle2 is a really refactored ContainsCycle1
func ContainsCycle2(head *Node) bool {

	tortoise := head
	hare := head

	for hare != nil {
		hare = hare.Next
		tortoise = tortoise.Next
		if hare != nil {
			hare = hare.Next
			if tortoise == hare {
				return true
			}
		}
	}
	return false
}

// CycleHead1 should find the node where a list intersects itself
// Does not find the correct node when given a circular list,
// a list without a head outside the cycle.
func CycleHead1(head *Node) *Node {

	tortoise := head
	hare := head
	curNode := head

	for {
		if hare == nil {
			return nil
		}
		hare = hare.Next
		if hare == nil {
			return nil
		}
		hare = hare.Next
		tortoise = tortoise.Next
		if hare == tortoise {
			break
		}
	}

	for {
		if tortoise == hare {
			tortoise = tortoise.Next
			curNode = curNode.Next
		}
		if tortoise == curNode {
			return curNode
		}
		tortoise = tortoise.Next
	}
}

// CycleHead2 should find the node where a list intersects itself
func CycleHead2(head *Node) *Node {

	tortoise := head
	hare := head

	for hare != nil {
		hare = hare.Next
		tortoise = tortoise.Next
		if hare != nil {
			hare = hare.Next
			if tortoise == hare {
				break
			}
		}
	}
	if hare == nil {
		return nil
	}
	hare = head
	for hare != tortoise {
		hare = hare.Next
		tortoise = tortoise.Next
	}
	return hare
}

// Floyds uses Floyd's algorithm to find if there's a cycle,
// at which link the cycle starts, and the cycle's length (period)
func Floyds(head *Node) (bool, *Node, int) {

	if head == nil {
		return false, nil, -1
	}

	// Setting hare and tortoise pointers are very important
	// for this algorithm to find meeting node.
	var hare *Node
	tortoise := head.Next
	if tortoise != nil {
		hare = tortoise.Next
	}

	// hare will get to nil first, if there's not a cycle
	for hare != nil && tortoise != hare {
		tortoise = tortoise.Next
		hare = hare.Next
		if hare != nil {
			hare = hare.Next
		}
	}

	if hare == nil || tortoise == nil {
		return false, nil, -1
	}

	// A cycle exists, don't have to check for end-of-list any more
	link := head
	tortoise = head
	for tortoise != hare {
		tortoise = tortoise.Next
		hare = hare.Next
		link = link.Next
	}

	period := 1
	hare = tortoise.Next

	for tortoise != hare {
		hare = hare.Next
		period++
	}

	return true, link, period
}
