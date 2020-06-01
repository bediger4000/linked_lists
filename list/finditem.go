package list

// FindItem returns a list node whose data matches the desired value,
// or nil if no node matches the desired value.
func FindItem(node *Node, desired int) *Node {
	for ; node != nil; node = node.Next {
		if node.Data == desired {
			return node
		}
	}
	return nil
}
