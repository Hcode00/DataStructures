package main

import (
	"testing"
)

func TestQueueAdd(t *testing.T) {
	queue := NewQueue[int]()

	// Add elements to the queue
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)

	// Check if the queue length is updated correctly
	if queue.GetLength() != 3 {
		t.Error("Queue length should be 3, but got", queue.GetLength())
	}

}
func TestQueuePeek(t *testing.T) {
	queue := NewQueue[int]()

	// Add elements to the queue
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)

	// Peek at the head value
	peekedValue := queue.Peek()

	// Check if the peeked value is correct
	if peekedValue != 1 {
		t.Error("Queue peeked value should be 1, but got", peekedValue)
	}

	// Check if the queue remains unchanged after peeking
	if queue.GetLength() != 3 {
		t.Error("Queue length should remain unchanged after peeking")
	}
}

func TestQueueIsEmpty(t *testing.T) {
	queue := NewQueue[int]()

	// Check if the empty queue is empty
	if !queue.IsEmpty() {
		t.Error("Empty queue should be empty")
	}

	// Add an element to the queue
	queue.Add(1)

	// Check if the non-empty queue is not empty
	if queue.IsEmpty() {
		t.Error("Non-empty queue should not be empty")
	}
}

func TestQueueGetLength(t *testing.T) {
	queue := NewQueue[int]()

	// Check the length of the empty queue
	if queue.GetLength() != 0 {
		t.Error("Empty queue length should be 0, but got", queue.GetLength())
	}

	// Add elements to the queue
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)

	// Check the length of the non-empty queue
	if queue.GetLength() != 3 {
		t.Error("Queue length should be 3, but got", queue.GetLength())
	}
}

func TestQueueContains(t *testing.T) {
	queue := NewQueue[int]()

	// Add elements to the queue
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)

	// Check if the queue contains element 2
	contains, index := queue.Contains(2)
	if !contains || index != 1 {
		t.Error("Queue should contain element 2 at index 1")
	}

	// Check if the queue does not contain element 4
	contains, index = queue.Contains(4)
	if contains || index != -1 {
		t.Error("Queue should not contain element 4")
	}
}

func TestQueuePop(t *testing.T) {
	queue := NewQueue[int]()

	// Add elements to the queue
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)

	// Pop an element from the queue
	poppedNode := queue.Pop()

	// Check if the popped element is correct
	if poppedNode.val != 1 {
		t.Error("Popped element value should be 1, but got", poppedNode.val)
	}

	// Check if the queue length is updated correctly
	if queue.GetLength() != 2 {
		t.Error("Queue length should be 2, but got", queue.GetLength())
	}
}

func TestQueuePrintData(t *testing.T) {
	queue := NewQueue[int]()

	// Add elements to the queue
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)

	// Print the contents of the queue
    t.Log("Queue contents:")
	queue.PrintData()

	// The expected output should be:
	// 1>>2>>3
    t.Log("Expected output: 1 -> 2 -> 3")
}
