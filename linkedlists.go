package main

type Data interface {
	int | string | float64
}

type Node[D Data] struct {
	val  D
	prev *Node[D]
	next *Node[D]
}

func newNode[D Data](val D, prev, next *Node[D]) *Node[D] {
	return &Node[D]{val: val, prev: prev, next: next}
}

type DoublyLinkedList[D Data] struct {
	head   *Node[D]
	tail   *Node[D]
	length int
}

// methods

func NewDoublyLinkedList[D Data]() *DoublyLinkedList[D] {
	return &DoublyLinkedList[D]{}
}

func (l *DoublyLinkedList[D]) Append(data D) {
	node := newNode[D](data, nil, nil)
	// case empty
	if l.head == nil || l.tail == nil {
		l.head = node
		l.tail = node
	}
	// case not empty
	if l.tail != nil {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
	l.length++
}

func (l *DoublyLinkedList[D]) Prepend(data D) {
	node := newNode[D](data, nil, nil)
	if l.tail == nil {
		l.tail = node
	}
	if l.head != nil {
		node.next = l.head
		l.head.prev = node
	}
	l.head = node
	l.length++
}

func (l DoublyLinkedList[D]) Get(data D) *Node[D] {
	node := l.head
	for {
		if node.val == data {
			return node
		}
		if node.next == nil {
			return nil
		}
		node = node.next
	}
}

func (l DoublyLinkedList[D]) GetLength() int {
	return l.length
}

func (l DoublyLinkedList[D]) IsEmpty() bool {
	if l.length == 0 {
		return true
	}
	return false
}

func (l DoublyLinkedList[D]) GetAtIndex(index int) *Node[D] {
	node := l.head
	count := 0
	for {
		if count == index {
			return node
		}
		if node.next == nil {
			return nil
		}
		count++
		node = node.next
	}
}

func (l DoublyLinkedList[D]) Contains(data D) (bool, int) {
	node := l.head
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

func (l *DoublyLinkedList[D]) InsertAt(data D, index int) *Node[D] {
	count := 0
	node := l.head

	// Handle head insertion
	if index == 0 {
		l.Prepend(data)
		return l.head
	}

	// Traverse the list
	for node != nil && count < index {
		node = node.next
		count++
	}

	// Handle end insertion (append)
	if node == nil {
		l.Append(data)
		return l.tail
	}

	// Handle middle insertion
	newNode := newNode[D](data, node.prev, node)
	node.prev.next = newNode
	node.prev = newNode

	l.length++
	return newNode
}

func (l *DoublyLinkedList[D]) RemoveAt(index int) *Node[D] {

	count := 0
	node := l.head
	if index < 0 || l.length == 0 {
		return nil
	}
	// Handle head deletion
	if index == 0 && l.length > 0 {
		sec := node.next
		sec.prev = nil
		l.head = sec
		node.next = nil
		l.length--
		return node
	}

	// Traverse the list
	for node != nil && count < index {
		node = node.next
		count++
	}

	// Handle bad index
	if node == nil {
		return nil
	}

	// Handle end deletion
	if node.next == nil {
		node.prev.next = nil
		l.tail = node.prev
		node.prev = nil
		l.length--
		return l.tail
	}

	// Handle middle deletion
	node.prev.next = node.next
	node.next.prev = node.prev
	l.length--

	return node

}

func (l *DoublyLinkedList[D]) Remove(data D) bool {
	node := l.head
	for {
		if node.val == data {
			if node.prev != nil && node.next != nil {
				node.prev.next = node.next
				node.next.prev = node.prev
				node.prev = nil
				node.next = nil
				l.length--
				return true
			}
			if node.prev == nil && node.next != nil {
				sec := node.next
				sec.prev = nil
				l.head = sec
				node.next = nil
				l.length--
				return true
			}

			if node.prev != nil && node.next == nil {
				node.prev.next = nil
				l.tail = node.prev
				node.prev = nil
				l.length--
				return true
			}
		}
		if node.next == nil {
			return false
		}
		node = node.next
	}
}
func (l DoublyLinkedList[D]) PrintData() {
	node := l.head
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


// general interface

type LinkedList[D Data] interface {
    PrintData()
    Contains(data D) (bool, int)
    GetLength() int
    IsEmpty() bool
}
