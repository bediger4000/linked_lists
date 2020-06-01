package list

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

// Really refactored ContainsCycle2
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
