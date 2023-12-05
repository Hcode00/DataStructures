package main

type Queue[D Data] struct {
	head   *QNode[D]
	tail   *QNode[D]
	length int
}

type QNode[D Data] struct {
    val  D
    next *QNode[D]
}

func newQNode[D Data](val D, next *QNode[D]) *QNode[D] {
    return &QNode[D]{val: val, next: next}
}

func NewQueue[D Data]() *Queue[D] {
	return &Queue[D]{}
}

func (q *Queue[D]) Add(data D) {
	node := newQNode[D](data, nil)
	if q.head == nil || q.tail == nil {
		q.head = node
		q.tail = node
	}
	if q.tail != nil {
        tail := q.tail
        tail.next = node
        q.tail = node
	}
	q.length++
}

func (q Queue[D]) IsEmpty() bool {
	if q.length == 0 {
		return true
	}
	return false
}

func (q Queue[D]) GetLength() int {
	return q.length
}

func (q Queue[D]) Contains(data D) (bool, int) {
	node := q.head
	count := 0
	for {
		if node.val == data {
			return true, count
		}
		if node.next == nil {
			return false, -1
		}
		count++
		node = node.next
	}
}

func (q *Queue[D]) Pop() *QNode[D] {
    head := q.head
    q.head = head.next
    head.next = nil
    q.length--
	return head
}

func (q Queue[D]) PrintData() {
	node := q.head
	count := 0
	if node == nil {
		println("Empty List !")
		return
	}
	for {
		if node.next == nil {
			println(node.val)
			return
		}
		print(node.val, " -> ")
		count++
		node = node.next
	}
}

func (q Queue[D]) Peek() D {
	return q.head.val
}
