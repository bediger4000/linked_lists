package heap

import "linked_lists/list"

// Delete removes the minimum-value item from the heap
// and returns the remaining heap and that min value
// item.
func (h Heap) Delete() (Heap, *list.Node) {
	// pop the next value from the top of the heap,
	// possibly leaving more linked list in place
	n := h[0]
	h[0] = n.Next

	// Here's the part that's different from a standard Min-Heap:
	//
	if h[0] == nil {
		h[0] = h[len(h)-1]
		h = h[:len(h)-1]
	}

	h.siftDown(0)

	return h, n
}

// siftDown checks for appropriate position of item
// in partially-ordered array at index idx.
// It moves items not in the right place.
func (h Heap) siftDown(idx int) {
	if idx > len(h)-1 {
		return
	}

	left := 2*idx + 1
	if left > len(h)-1 {
		return
	}
	if h[idx].Data > h[left].Data {
		h[idx], h[left] = h[left], h[idx]
		h.siftDown(left)
	}

	right := 2*idx + 2
	if right > len(h)-1 {
		return
	}
	if h[idx].Data > h[right].Data {
		h[idx], h[right] = h[right], h[idx]
		h.siftDown(right)
	}
}
