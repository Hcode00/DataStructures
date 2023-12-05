package main

type Stack[D Data] struct {
	head   *SNode[D]
	tail   *SNode[D]
	length int
}

type SNode[D Data] struct {
	prev *SNode[D]
	val  D
}

func newSNode[D Data](val D, prev *SNode[D]) *SNode[D] {
	return &SNode[D]{val: val, prev: prev}
}

func NewStack[D Data]() *Stack[D] {
	return &Stack[D]{}
}

func (q *Stack[D]) Add(data D) {
	node := newSNode[D](data, nil)
	if q.head == nil || q.tail == nil {
		q.head = node
		q.tail = node
	}
	if q.tail != nil {
		tail := q.tail
		node.prev = tail
		q.tail = node
		q.head.prev = nil
	}
	q.length++
}

func (q Stack[D]) IsEmpty() bool {
	if q.length == 0 {
		return true
	}
	return false
}

func (q Stack[D]) GetLength() int {
	return q.length
}

func (q Stack[D]) Contains(data D) (bool, int) {
	node := q.tail
	count := q.length - 1
	for {
		if node.val == data {
			return true, count
		}
		if count == 0 {
			return false, -1
		}
		count--
		node = node.prev
	}
}

func (q *Stack[D]) Pop() *SNode[D] {
	tail := q.tail
	q.tail = tail.prev
	q.length--
	return tail
}

func (q Stack[D]) PrintData() {
	node := q.tail
	count := q.length - 1
	if node == nil {
		println("Empty List !")
		return
	}
	for {
		if node.prev == nil || count == 0 {
			println(node.val)
			return
		}
		print(node.val, " -> ")
		count--
		node = node.prev
	}
}

func (q Stack[D]) Peek() D {
	return q.tail.val
}
