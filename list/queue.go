package list

// Enqueue puts value on the end of the queue.
func (q *Queue) Enqueue(value int) {
	in := &Node{Data: value}
	in.Next = q.in
	q.in = in
}

// Dequeue returns the value of the first item put on
// the queue. That value is no longer on the queue,
// Dequeue is destructive
func (q *Queue) Dequeue() int {
	if q.Empty() {
		return 0
	}
	if q.out == nil {
		in := &(q.in)
		out := &(q.out)
		for !Empty(in) {
			Push(out, Pop(in))
		}
	}
	return Pop(&(q.out))
}

// Empty returns true when called with a queue
// that has no elements.
func (q *Queue) Empty() bool {
	if q.out != nil {
		return false
	}
	if q.in != nil {
		return false
	}
	return true
}
