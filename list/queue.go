package list

func (q *Queue) Enqueue(value int) {
	in := &Node{Data: value}
	in.Next = q.in
	q.in = in
}

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

func (q *Queue) Empty() bool {
	if q.out != nil {
		return false
	}
	if q.in != nil {
		return false
	}
	return true
}
